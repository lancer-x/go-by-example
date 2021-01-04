package main

import "fmt"

func main()  {
	a := []int{0,0,0,0,1,1,1,2,3,3,4,4,4,4,4,4,6,7}
	solution(a)
}

func solution(a []int) int {
	for i := 1; i < len(a); i++ {
		if a[i] == a[i-1] {
			a = append(a[:i-1], a[i:]...)
			//因为当前元素向后覆盖了一位，所以同事i要相应的减1
			i--
		}
	}

	fmt.Println(a)
	return len(a)
}