package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch1 = make(chan string, 5)

//var ch2 = make(chan string)

func main()  {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			fmt.Println("Ticker tick.")
		}
	}()
	n := 10
	for i:=0; i<n; i++ {
		go test(ch1)
	}

	for i:=0; i<n; i++ {
		s := <-ch1
		fmt.Println(s)
	}
}

func test(ch chan string)  {
	t := randInt()
	time.Sleep(time.Duration(t) * time.Second)
	msg := fmt.Sprintf("time -- %d", t)
	ch <- msg
}

func testBuffer(ch chan string)  {

}

func randInt() int {
	t1 := time.Now().UnixNano()

	//cTimeStamp := t1.UnixNano()

	rand.Seed(t1)

	var a = rand.Intn(10)
	return a
}
