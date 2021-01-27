package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println(m)
			time.Sleep(1 * time.Second)
		}
	}()
	ch <- "aaa"
	ch <- "bbb"
	ch <- "ccc"
}
