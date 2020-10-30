package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func main()  {
    var wg sync.WaitGroup
    num := 100
    wg.Add(num)
    for i := 0; i < num; i++ {
        go run(&wg)
    }
    wg.Wait()
}

func run(wg *sync.WaitGroup)  {
    defer wg.Done()
    rand.Seed(time.Now().Unix())
    seconds := rand.Intn(10)
    fmt.Println(seconds)

}
