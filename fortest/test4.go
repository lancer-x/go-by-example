package main

import "strings"

func main() {
	a := "abcabc"
	c := strings.Split(a, "c")
}

func bonus(totalMoney float32, num int) []float32 {
    bonus := make([]float32, num)
    for num >= 1 {
        nowTotalMoney, perMoney = getMoney(totalMoney, num)
        if num == 1 {
            bonus = append(bonus, round(nowTotalMoney, 2))
        } else {
            bonus = append(bonus, round(perMoney, 2))
        }
        num--
    }
}

func getMoney(money float64, num int) float32, float32 {
    perMoney := money / num * random(0.5, 1.5)
    return money - perMoney, perMoney
} 

/*
file
统计输出每个司机评价次数最多的前五条评论
//id,comment 
*/



