package main

import "fmt"

func main() {
    t1()
}

func t1() {
	//直接初始化

	s1 := []int{1, 2, 3}

	//使用make
	s2 := make([]int, 3)
	s2[0] = 3

    //仅声明的slice可以直接append
	var s3 []int
    fmt.Printf("%t\n", s3 == nil)
    s3 = append(s3, 1)
    fmt.Printf("%t\n", s3 == nil)

	fmt.Println(s1, s2, s3)
}


