package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for true {
			select {
			case <-ctx.Done():
				fmt.Println("正在退出")
				//runtime.Goexit()
				return
			default:
				fmt.Println("正在工作")
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	time.Sleep(time.Second)
}
