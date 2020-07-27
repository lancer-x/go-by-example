package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s);  {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}

	byte2string()
}

func byte2string()  {
	s := "abcdef"
	b := []byte(s)
	s2 := string(b)

	s = "aaaa"
	b[2] = '2'
	fmt.Println(s, b, s2)
}

