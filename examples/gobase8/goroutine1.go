package main

import (
	"fmt"
	"time"
)

func main1()  {
	go spinner(100 * time.Millisecond)
	fibR := fib(45)
	fmt.Println(fibR)
}

func spinner(delay time.Duration)  {
	for {
		for _,r := range `-\|/` {
			fmt.Print("-")
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}