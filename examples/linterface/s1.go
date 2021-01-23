package main

import "fmt"

type People interface {
	Speak(word string) string
}

type student struct {

}

func (s *student) Speak(word string) string {
	if word == "aaa" {
		fmt.Println("aaa")
	} else {
		fmt.Println("bbb")
	}
	return "aa"
}

func main()  {
	var stu People = &student{}
	stu.Speak("bb")
}
