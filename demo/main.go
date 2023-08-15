package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()

		for {

			select {
			case <-ctx.Done():
				fmt.Println("done")
				goto done
			default:
				fmt.Println("continue")
			}
			time.Sleep(2 * time.Second)
		}
	done:
		return
	}(ctx)

	wg.Wait()
}
