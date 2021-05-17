package main

import "time"

func main() {
	mainChan := handle()
	evenChan := printEven(mainChan)
	go printOdd(evenChan, mainChan)
	time.Sleep(time.Second)

}

func handle() chan int {
	ch := make(chan int)
    go func ()  {
        ch <- 1
    }()
    
	return ch
}

func printEven(ch <-chan int) chan int {
	evenChan := make(chan int)
	go func() {
		for i := range ch {
			evenChan <- i 
			println("even",i * 2)
		}
	}()
	return evenChan
}

func printOdd(ch <-chan int, mainChan chan int) {
	for i := range ch {
		println("odd", i * 2 - 1)
        if ok := doNext(mainChan, i +1, 100); !ok {
            return
        }
	}
}

func doNext(ch chan int, num,target int) bool {
    if num > target {
        close(ch)
        return false
    }
    ch <- num
    return true
}
