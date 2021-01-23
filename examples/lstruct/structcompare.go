package main

import "fmt"

func main()  {
	s1 := struct {
		age int
		name string
	}{age: 111, name: "qq"}

	s2 := struct {
		age int
		name string
	}{age: 111, name: "qq"}

	if s1 == s2 {
		fmt.Println("equal")
	}

	s3 := struct {
		name string
		age int
	}{name: "qq", age: 11}

	if s1 == s3 {
		fmt.Println("equ")
	}
}
