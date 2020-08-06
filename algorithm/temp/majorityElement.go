package main

import (
	"fmt"
	"math"
)

func main()  {
	nums := []int{2,2,1,1,1,2,2}
	a := majorityElement(nums)
	fmt.Println(a)
}

func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	limit := int(math.Ceil(float64(len(nums)) / 2))
	for _,v := range nums {
		l := addNum(numMap, v)
		if l >= limit {
			return v
		}
	}
	return 0
}

func addNum(numMap map[int]int, key int) int {
	numMap[key]++
	return numMap[key]
}
