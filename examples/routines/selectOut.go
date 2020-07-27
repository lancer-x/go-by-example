package main

import (
	"fmt"
	//"os"
	"time"
)

func main()  {
	channel := make(chan bool)
	ch2 := make(chan int)
	go worker(channel, ch2)
	time.Sleep(time.Second)
	channel <- true
	<- ch2
}

func worker(channel chan bool, ch chan int)  {
	for {
		select {
		default:
			fmt.Println("hello")
		case <- channel:
			fmt.Println("noono")
			ch <- 1
		}
	}
}