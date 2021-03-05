package main

import "fmt"

func reverseString(s []byte)  {
	str := []byte{'H','a','n','n','a','h'}
	reverse(str, 0, len(str) - 1)
	fmt.Println(str)
}

func reverse(s []byte, start int, end int) {
	if start >= end {
		return
	}
	s[start],s[end] = s[end],s[start]
	reverse(s, start + 1, end - 1)
}