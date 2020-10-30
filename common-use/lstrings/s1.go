package main

import (
	"fmt"
	"strings"
)

func main()  {
	tEqualFold()
}

//判断两个字符串(将大写，小写，标题格式字符视为相同)是否相同
func tEqualFold() {
	a := "abcdefs"
	b := "Abcdefs"
	fmt.Println(strings.EqualFold(a, b))
}

//判断s是否有前缀