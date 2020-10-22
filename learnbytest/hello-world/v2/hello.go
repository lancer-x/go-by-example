package main

import "fmt"

func main()  {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	return "hello," + name
}