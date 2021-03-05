package main

import (
	"fmt"
	"time"
)

func main2()  {
	stream := make(chan string)
	stream2 := make(chan string)
	go service1(stream)
	go service2(stream2)

	select {
	case t := <-stream:
		fmt.Println(t)
	case t := <-stream2:
		fmt.Println(t)
	case <-time.After(10 * time.Second):
		fmt.Println("time up")
	}
}

func service1(ch chan string)  {
	time.Sleep(3 * time.Second)
	ch <- "services1 run"
}


func service2(ch chan string)  {
	time.Sleep(5 * time.Second)
	ch <- "services2 run"
}