package main

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	for {
		endpoints, err := clientset.CoreV1().Endpoints("ohmyfans").Get(context.Background(), "subs-rpc-svc", metav1.GetOptions{})
		fmt.Println(err)
		fmt.Println(endpoints)
		time.Sleep(2 * time.Second)
	}

}
