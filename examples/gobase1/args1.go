package main

import (
	"fmt"
	"os"
	"strings"
)

func main2()  {
	var s, sep string

	for i := 1; i< len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	map1 := getArgsMap()
	fmt.Println(map1)
}

func getArgsMap() map[string]interface{} {
	argsMap := make(map[string]interface{})

	for i := 1; i < len(os.Args); i++ {
		if !strings.Contains(os.Args[i], "=") {
			argsMap[os.Args[i]] = ""
			continue
		}
		tmp := strings.Split(os.Args[i], "=")

		argsMap[tmp[0]] = tmp[1]
	}
	return argsMap
 }