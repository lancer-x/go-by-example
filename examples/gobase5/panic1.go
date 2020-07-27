package main

import (
	"errors"
	"fmt"
)

func main()  {
	//panic1()
	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("sth is wrong")
			fmt.Println(err)
		}
	}()
	recover1()
}

func panic1()  {
	err := errors.New("this is panic")
	panic(err)
}

func recover1() (err error) {
	panic1()

	return
}