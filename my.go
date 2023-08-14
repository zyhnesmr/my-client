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
		res, err := clientset.CoreV1().Services("ohmyfans").Get(context.Background(), "subs-rpc-svc", metav1.GetOptions{})
		fmt.Println(err)
		fmt.Println(res)
		time.Sleep(time.Second * 2)
	}

}
