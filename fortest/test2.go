package main

import "fmt"

func main2() {
    lmap := make(map[string]string)

    lmap["aaa"] = "aaa";
    lmap["bbb"] = "bbb";
    lmap["ccc"] = "ccc";

    fmt.Println(lmap)

    for k, v := range lmap {
        fmt.Print(k)
        v = v + "1"
    }
    fmt.Println(lmap)
}