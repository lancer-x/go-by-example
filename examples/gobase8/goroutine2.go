package main

import (
	"fmt"
	//"time"
)

var ch = make(chan int)

var ch2 = make(chan int)

var ch3 = make(<-chan int)

func main()  {
	go func() {
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
			//time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			i,ok := <-ch
			if (!ok) {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	for  {
		i,ok := <- ch2
		if (!ok) {
			fmt.Println("all channels are closed")
			break
		}
		fmt.Println(i)
	}
}