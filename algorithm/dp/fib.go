package main

import "fmt"

func main()  {
	a := recursiveSolution(10)
	fmt.Println(a)
}

//纯递归解法
func recursiveSolution(n int) int{
	fmt.Println("now", n)
	if n <= 1 {
		return n
	}

	return recursiveSolution(n - 1) + recursiveSolution(n - 2)
}