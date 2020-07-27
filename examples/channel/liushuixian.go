package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := producer(1,2,3,4,5)
	ch2 := handleNum1(ch1)
	ch3 := handleNum2(ch2)
	flag := show(ch3)
	<-flag
}


//生产原始数据
func producer(nums ...int) <- chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _,i := range nums {
			fmt.Printf("原始数据生产完成%d\n", i)
			time.Sleep(time.Second)
			ch <- i
		}
	}()
	return ch
}

func handleNum1(chin <-chan int) <- chan int {
	outch := make(chan int)

	go func() {
		defer close(outch)
		for i := range chin {
			fmt.Printf("hander 1获取数据%d\n", i)
			outch <- i + 1
		}
	}()
	return outch
}

func handleNum2(chin <-chan int) <- chan int{
	outch := make(chan int)
	go func() {
		defer close(outch)
		for i := range chin {
			fmt.Printf("hander 2 获取数据%d\n", i)
			outch <- i * i
		}
	}()
	return outch
}

func show(chin <-chan int) chan bool{
	ch := make(chan bool)
	go func() {
		defer func() {
			ch <- true
		}()
		for i := range chin {
			fmt.Printf("处理完成%d\n-----------------\n\n", i)
		}
	}()
	return ch
}