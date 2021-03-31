package main

import "fmt"

type NewInt int //

type IntAlias = int

func main() {
	var a NewInt //
	var b IntAlias
	fmt.Printf("%T %T", a, b)
}