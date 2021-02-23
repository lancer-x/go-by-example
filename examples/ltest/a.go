package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ch := make(chan string, 1)
	var res strings.Builder

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			res.WriteString(".")
			res.WriteString(x)
		case ch <- string(rune(i)):
			res.WriteString("-")
			res.WriteString(strconv.Itoa(i))
			//fmt.Println(i)
		}
	}
	fmt.Println(res.String())
}

/*
0  ch空

*/
