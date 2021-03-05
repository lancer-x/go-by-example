package main

import (
	"fmt"
	"log"
	"time"
)

func main1()  {
	b := doadd(5, addone)
	fmt.Println(b)

	bigSlowOperation()
}

func addone(v int) int {
	return v + 1
}

func doadd(v int, f func(int) int) int {
	return f(v)
}


func bigSlowOperation() {
	defer trace("kkkkkkkk")()

	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}