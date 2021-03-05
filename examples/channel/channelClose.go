package main

import "fmt"

func main1()  {
	c := make(chan int)
	close(c)

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}
