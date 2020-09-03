package main

import "fmt"

func main()  {
	test()
}

func test()  {
	s := "hello 瓜子"
	for i := 0 ; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for k,v := range s {
		fmt.Println(k, string(v))
	}

	rr := []rune(s)
	fmt.Println(len(rr), string(rr[7]))

}