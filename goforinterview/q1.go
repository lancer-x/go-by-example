package main

import "fmt"

func main() {
    //test1()
    test2()
}

func test1 () {
    defer func() {
        err := recover()
        switch err.(type) {
        case error:
            fmt.Println(err)
        default:
            fmt.Println("default")
        }

    }()
    defer func() {fmt.Println("打印前")}()
    defer func() {fmt.Println("打印中")}()
    defer func() {fmt.Println("打印后")}()
    panic(fmt.Errorf("打印完成"))
}

func test2() {
    a := 5
    b := 8.1
    fmt.Println(a + b)
}