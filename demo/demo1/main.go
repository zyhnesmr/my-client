package main

import (
	"fmt"
	"time"
)

func main() {

	var c = make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println(<-c)
	}()

	c <- true
	time.After(time.Second * 10)
}
