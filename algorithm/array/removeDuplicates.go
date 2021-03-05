package main

func main3() {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(a)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for i < len(nums) {
		j := i + 1
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		if j-1 >= i {
			nums = append(nums[:i+1], nums[j:]...)
		}
		i++
	}
	return len(nums)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	helper(nums, 0, k-1)
	helper(nums, k, len(nums)-1)
	helper(nums, 0, len(nums)-1)
}

func helper(nums []int, i int, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
