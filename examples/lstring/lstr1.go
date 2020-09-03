package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "this is 中国"
	fmt.Println(len(s))

	size := utf8.RuneCountInString(s)

	fmt.Println(size)
}
