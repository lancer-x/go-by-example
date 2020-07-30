package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//主动取消
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	go childDo(ctx, "任务1")
	go childDo(ctx, "任务2")
	go childDo(ctx, "任务3")
	defer cancel()
	time.Sleep(15 * time.Second)

}


func childDo(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("task is done")
			return
		default:
			fmt.Println(name, " is runnning")
			time.Sleep(2 * time.Second)
		}
	}
}