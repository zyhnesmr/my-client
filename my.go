package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go Monitor(clientset, &wg)

	wg.Wait()
}

func Monitor(clientset *kubernetes.Clientset, wg *sync.WaitGroup) {
	defer wg.Done()
	watcher, err := clientset.CoreV1().Endpoints("ohmyfans").Watch(context.Background(), metav1.ListOptions{
		FieldSelector: "metadata.name=" + "subs-rpc-svc",
	})

	if err != nil {
		panic(err)
	}

	for {
		select {
		case event, ok := <-watcher.ResultChan():
			if !ok {
				fmt.Println("Endpoints watcher channel not ok")
				fmt.Println(event)
				fmt.Println("-----new watcher-------")

				for {
					watcher, err = clientset.CoreV1().Endpoints("ohmyfans").Watch(context.Background(), metav1.ListOptions{
						FieldSelector: "metadata.name=" + "subs-rpc-svc",
					})
					if err == nil {
						break
					}
					time.Sleep(time.Second * 1)
				}
				fmt.Println("regen watcher")
			} else {
				eds, ok := event.Object.(*corev1.Endpoints)
				if !ok {
					fmt.Println("not endpoints")
				}

				switch event.Type {
				case watch.Added:
					fmt.Println("---------WatchAdded-----------")
					for _, subset := range eds.Subsets {
						for _, address := range subset.Addresses {
							fmt.Println(address.IP)
						}
					}
				case watch.Modified:
					fmt.Println("---------WatchModified-----------")
					for _, subset := range eds.Subsets {
						for _, address := range subset.Addresses {
							fmt.Println(address.IP)
						}
					}
				case watch.Deleted:
					fmt.Println("---------WatchDeleted-----------")
					for _, subset := range eds.Subsets {
						for _, address := range subset.Addresses {
							fmt.Println(address.IP)
						}
					}
				}
			}
		}
	}
}
