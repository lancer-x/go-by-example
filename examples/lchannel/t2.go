package main

import (
	"fmt"
	"time"
)

func main3() {
	var ch chan int
	select {
	case <-ch:
		fmt.Println("ch is doing the work")
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
	fmt.Println("done")
}
