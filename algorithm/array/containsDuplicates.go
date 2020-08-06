package main

func containsDuplicate(nums []int) bool {
	lmap := make(map[int]int)

	for _,v := range nums {
		lmap[v]++
		if (lmap[v] > 1) {
			return true
		}
	}
	return false
}


