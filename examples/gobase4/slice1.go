package main

import "fmt"

func main()  {
	//slice 不可以直接进行比较，butes.Equal可以对byte切片进行比较
	//唯一合法的比较操作是和nil
	//summer := []string{}

	//summer1 := make([]string, 2)

	//slice判空使用 len(summer) == 0,因为nil为0，空切片也是0

	var summer3 []string

	if summer3 == nil {
		fmt.Println("it's nil", len(summer3))

	}
}



func  reverse(s []string){
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func mergeNearSameChar(s []string) {

}