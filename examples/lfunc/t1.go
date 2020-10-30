package main

import (
	"fmt"
	"strings"
)

func getSumFunc(c int) func(int, int) int {
	return func(x int, y int) int {
		return x + y + c
	}
}

func main()  {
	a := getSumFunc(3)
	b := a(1,2)
	println(b)

	println(doAdd(5, 6,  addNum))

	f1(helpf(f2, 100, 200))

	testAddSuffixFunc()
}

func doAdd(x int, y int, add func(int, int) int) int {
	return add(x, y)
}

func addNum(x int, y int) int {
	return x + y
}


//f1 f2 两个函数已确定，实现将f2当作参数传至f1

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int)  {
	fmt.Println("this iis f2")
	fmt.Println(x + y)
}

func helpf(f func(x, y int), x, y int) func() {
	return func() {
		f(x, y)
	}
}

//给字符串加不同的结尾
func addSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}

func testAddSuffixFunc()  {
	jpegFunc := addSuffixFunc(".jpeg")
	gifFunc := addSuffixFunc(".gif")

	a := jpegFunc("aaaa")
	b := gifFunc("bbbb")
	fmt.Println(a, b)
}