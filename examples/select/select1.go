package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	i1, i2 := 2, 3

	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <-ch1:
				fmt.Println("ch1 out")
			case ch1 <- i1:
				fmt.Println("ch1 in", i1)
			case <-ch2:
				fmt.Println("ch2 out")
			default:
				fmt.Println("default")
			}
		}
	}()

	ch1 <- i1
	ch1 <- i1
	ch1 <- i1

	<-ch1

	i2++
	time.Sleep(10 * time.Second)
	fmt.Println("done")
}
