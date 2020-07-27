package main

import (
	"fmt"
	"flag"
	"strings"
)

var n = flag.Bool("n", false, "去除回车")
var sep = flag.String("s", " ", "分隔符")

func main()  {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println("----")
	}
}
