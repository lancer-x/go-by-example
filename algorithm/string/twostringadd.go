package main

import "fmt"

func main() {
	s1 := "12345"
	fmt.Print()
}

func doAdd(str1, str2 string) string {
	if len(str1) == 0 {
        return str2
    }
    if len(str2) == 0 {
        return str1
    }
    

}