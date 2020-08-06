package main

import "sort"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	res := make([]int,2)

	for i,v := range nums {
		if other,ok := numMap[target-v]; ok && i != other{
			res[0],res[1] = i, other
			break
		}
		numMap[v] = i
	}
	return res
}


func isAnagram(s string, t string) bool {
	sort.Slice([]byte(s), func(i, j int) bool {
		return s[i] > s[j]
	})

	sort.Slice([]byte(t), func(i, j int) bool {
		return t[i] > t[j]
	})

	return s == t
}