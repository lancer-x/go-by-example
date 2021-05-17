package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
    var a int = 0
    threads := runtime.GOMAXPROCS(0)

    for i:=0; i < threads; i++ {
        go func () {
            a++
        }()
    }
    time.Sleep(time.Second)
    fmt.Print(a)
}