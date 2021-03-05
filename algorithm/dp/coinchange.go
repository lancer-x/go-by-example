package main

import "math"
import "fmt"

var memo = make(map[int]int)
func main1()  {
	coins := []int{2,5,7}
	//memo[0] = 0
	/*for v := range coins {
		memo[v] = 1
	}*/
	print(solution(27, coins))
}

func solution(money int, coins []int) int {
	if money <=0 || len(coins) == 0 {
		return 0
	}
	return dp(money, coins)
}

func dp(money int, coins []int) int {
	if v,ok := memo[money]; ok {
		return v
	}
	if money == 0 {
		return 0
	}
	if money < 0 {
		return -1
	}


	res := math.MaxInt32
	for _, coin := range coins {
		subProblem := dp(money - coin, coins)
		if subProblem == -1 {
			continue
		}
		res = min(subProblem + 1, res)
	}

	if res != math.MaxInt32 {
		memo[money] = res
	} else {
		memo[money] = -1
	}
	fmt.Println(memo)

	return memo[money]
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}