package main

func main2() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
}

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					swap(nums, i, j)
					break
				}
			}
		}
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
