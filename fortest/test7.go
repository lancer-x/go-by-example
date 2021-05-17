package main

import "fmt"

func main() {
	arr := []int{2, 3, 1, 6, 5, 4}
	heapSort(arr)
	fmt.Println(arr)
}

func heapSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertHeap(arr, i)
        fmt.Println(arr)
	}
}

func insertHeap(arr []int, i int) {
    fmt.Println(arr[i])
	for arr[i] > arr[i/2] {
        fmt.Println("--", i, i/2)
        arr[i], arr[i/2] = arr[i/2], arr[i]
		i = i / 2
	}
}