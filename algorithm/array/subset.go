package main

func main()  {

}

func do(nums []int) [][]int {
	rmap := make(map[int]int)
	ret := make([][]int, 0)
	for i:=0; i<len(nums); i++ {
		if _,ok := rmap[nums[i]]; ok {
			continue
		}
		l := len(ret)
		for j := 0; j < l; j++ {
			ret = append(ret, append(ret[j], nums[i]))
		}
		ret = append(ret, []int{nums[i]})
	}
	return append(ret, []int{})
}