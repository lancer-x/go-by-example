package main

import (
	"fmt"
)

func main()  {
	result := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 20; i++ {
				ch <- i
			}
		}()
		return ch
	}

	resultChan := result()
	for i := range resultChan {
		fmt.Println(i)
	}
}


