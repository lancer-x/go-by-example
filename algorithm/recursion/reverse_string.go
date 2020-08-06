package main

import "fmt"

func main()  {
	s := "abcdefghijk"
	reverseString([]byte(s))
}

func reverseString(s []byte) {
	res := make([]byte, 0)
	reverse(s, 0, &res)
	for i := 0; i < len(s); i++ {
		s[i] = res[i]
	}

	fmt.Println(string(s))
}
func reverse(s []byte, i int, res *[]byte) {
	if i == len(s) {
		return
	}
	reverse(s, i+1, res)
	*res = append(*res, s[i])
	fmt.Println(i, string(s[i]), "+", string(*res))
}


