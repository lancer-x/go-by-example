package main

import "fmt"

func main() {
    t1()
}

func t1() {
	count := 100
	srcData := make([]int, count)
	for i := 0; i < count; i++ {
		srcData[i] = i
	}

	refData := srcData

	dstData := make([]int, count)

	copy(dstData, srcData)

	srcData[0] = 1111

	fmt.Printf("%d %d %d\n", srcData[0], refData[0], dstData[0])

    copy(dstData, srcData[:5])

    for i := 0; i < 10; i++ {
        fmt.Println(dstData[i])
    }

    fmt.Printf("%d %d", len(dstData), cap(dstData))

}