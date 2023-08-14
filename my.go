package main

import (
	"context"
	"fmt"

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

	//endpoints, err := clientset.CoreV1().Endpoints("ohmyfans").Get(context.Background(), "subs-rpc-svc", metav1.GetOptions{})
	//fmt.Println(err)
	//fmt.Println(endpoints)
	//time.Sleep(2 * time.Second)
	//
	//for _,subset:=range endpoints.Subsets{
	//	for _,address:=range subset.Addresses{
	//		fmt.Println(address.IP)
	//	}
	//}

	watcher, err := clientset.CoreV1().Endpoints("ohmyfans").Watch(context.Background(), metav1.ListOptions{
		FieldSelector: "metadata.name=" + "subs-rpc-svc",
	})
	fmt.Println(err)

	go func() {
		for {
			select {
			case event, ok := <-watcher.ResultChan():
				if !ok {
					fmt.Println("Endpoints watcher channel closed.")
					return
				}

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
				default:
					break
				}
			}
		}
	}()
}
