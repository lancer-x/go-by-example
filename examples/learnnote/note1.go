package main

import (
	"fmt"
	"strconv"
)

func main()  {
	a,b := getValue()
	fmt.Println(a,b)
	fmt.Println(test())
}

func t()  {
	a := "111"
	b, _ := strconv.ParseInt(a, 10, 0)
	c, _ := strconv.Atoi(a)
	fmt.Println(b, c)

	d := strconv.FormatInt(b, 10)
	fmt.Println(d)
	f := 23432353.124342342532553
	e := strconv.FormatFloat(f, 1, 2, 32)
	fmt.Println(e)
}

func getValue()  (a int, b int){
	//a,b := 1, 2
	defer func() {
		a += 1
		b += 1
		fmt.Println(a,b)
	}()
	return 1,2
}

func test() (z int) {
	defer func() {
		z += 100
	}()
	return 100
}

