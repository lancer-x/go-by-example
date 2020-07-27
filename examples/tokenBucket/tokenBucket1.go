package main

import (
	"fmt"
	"time"
)

func main()  {
	var fillTime = time.Millisecond * 10
	var cap = 100
	var tokenBucket = make(chan struct{}, cap)

	go func() {
		ticker := time.NewTicker(fillTime)
		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:

				}
				fmt.Println("当前数量", len(tokenBucket), time.Now())
			}
		}
	}()
	time.Sleep(time.Second)
}

