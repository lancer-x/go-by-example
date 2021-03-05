package main

import (
	"fmt"
)

//pirnt fish cat dog
func main3() {
	sigChan := make(chan int)
	fishChan := make(chan int)
	catChan := make(chan int)
	dogChan := make(chan int)

	fishTask := task()
	catTask := task()
	dogTask := task()

	go fishTask("fish", fishChan, catChan, sigChan, false)
	go catTask("cat", catChan, dogChan, sigChan, false)
	go dogTask("dog", dogChan, fishChan, sigChan, true)

	fishChan <- 1
	var counter int
	for range sigChan {
		counter++
		fmt.Println(counter)
		fishChan <- 1
		if counter == 100 {
			close(sigChan)
			break
		}

	}
}

func task() func(string, chan int, chan int, chan int, bool) {
	return func(name string, ch, nextCh, sig chan int, lastTask bool) {

		for i := range ch {
			fmt.Print(name)
			if lastTask {
				sig <- 0
			} else {
				nextCh <- i
			}

		}
	}
}
