package main

import (
	"fmt"
)

func main()  {
	//a := []int{5,7,6,8,1,4,16,3}
	//a := []int{1,3,4,5}
	//a := []int{8,6,16,7}
	//a := []int{5,2,3,1}
	a := []int{}
	fmt.Println(sortArray(a))
}

func sortArray(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}

	 quicksort(nums, 0, len(nums) - 1)
	return nums
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		quickSort(arr, left, partitionIndex - 1)
		quickSort(arr, partitionIndex + 1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index - 1] = arr[index - 1], arr[pivot]
	return index - 1
}



func quicksort(nums []int, left int, right int) {
	if left < right {
		i,j := left+1,right

		baseNum := nums[left]

		for  {
			for nums[j] > baseNum {
				//从后向前找到第一个小于base的元素
					j--
			}
			for nums[i] < baseNum && i < len(nums) - 1{
				//从前向后找到第一个大于base的元素
					i++

			}
			//fmt.Println(nums)
			if i < j {
				nums[i],nums[j] = nums[j],nums[i]
			} else {
				break
			}
			//time.Sleep(time.Second)
		}

		//fmt.Println(nums, i, j)

		nums[left],nums[j] = nums[j], baseNum
		quicksort(nums, left, j-1)
		quicksort(nums, j + 1, right)
	}

}


