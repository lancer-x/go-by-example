package main

import (
	"fmt"
	"time"
)

func main()  {
	flag := make(chan int)
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	go tickDo(t)

	tch := time.After(10 * time.Second)

	go delayDo(tch, flag)
	<-flag
}

func tickDo(t *time.Ticker)  {
	for  {
		select {
		case <-t.C:
			fmt.Println("exec the task")
		default:
			time.Sleep(time.Second)
		}
	}
}

func delayDo(ch <-chan time.Time, flag chan int)  {
	for {
		select {
		case <-ch:
			fmt.Println("time is coming")
			flag <- 1
			return
		default:
			fmt.Println("waiting")
			time.Sleep(time.Second)
		}
	}
}