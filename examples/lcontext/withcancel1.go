package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "任务1")
	go watch(ctx, "任务2")
	go watch(ctx, "任务3")
	time.Sleep(10 * time.Second)

	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "退出")
			return
		default:
			fmt.Println(name, "监控中")
			time.Sleep(2 * time.Second)
		}
	}
}