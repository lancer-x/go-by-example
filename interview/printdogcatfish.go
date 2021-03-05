package main

import (
	"fmt"
)

func main2() {
	sigChan := make(chan int)
	catChan := make(chan int)
	dogChan := make(chan int)
	fishChan := make(chan int)
	go printCat(catChan, dogChan)
	go printDog(dogChan, fishChan)
	go printFish(fishChan, catChan, sigChan, 100)
	catChan <- 1

	// for i := 0; i < 100; i++ {
	// 	catChan <- i
	// 	fmt.Println(i)
	// }

	//time.Sleep(1 * time.Second)
	<-sigChan
}

func printCat(ch <-chan int, nextCh chan<- int) {
	for i := range ch {
		fmt.Print("cat")
		nextCh <- i
	}
}

func printDog(ch <-chan int, nextCh chan<- int) {
	for i := range ch {
		fmt.Print("dog")
		nextCh <- i
	}
}

func printFish(ch <-chan int, nextCh chan<- int, sigChan chan int, count int) {
	for i := range ch {
		fmt.Print("fish\n")
		if i < count {
			i++
			nextCh <- i
		} else {
			sigChan <- 1
		}
	}
}
