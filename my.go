package main

import (
	"context"
	"fmt"
	_ "my-name-resolver/resolvers"
	"my-name-resolver/subs"
	"time"

	"google.golang.org/grpc"
)

func main() {

	clientConn, err := grpc.Dial("omf://ohmyfans/subs-rpc-svc:9091", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	subscribeClient := subs.NewSubscribeClient(clientConn)

	for {
		getGroupList, err := subscribeClient.GetGroupList(context.TODO(), &subs.UserId{Id: 1})
		fmt.Println(err)
		fmt.Println(getGroupList)

		time.Sleep(time.Second * 2)
	}
}
