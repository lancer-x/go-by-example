package main

import "fmt"

func main()  {
	a := make([]int, 0)
	for i := 0; i < 20; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	reverse(a)
}

func reverse(a []int) {
	l := len(a)

	for i := 0; i < l - i - 1; i++ {
		a[i], a[l-i - 1] = a [l-i - 1], a[i]
	}
	fmt.Println(a)
}
