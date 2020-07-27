package main

import (
	"fmt"
	"reflect"
)

func main()  {
	t := reflect.TypeOf(3)
	fmt.Println(t)



	v := reflect.ValueOf(3)
	fmt.Println(v)

	T := v.Type()
	fmt.Println(T.String())
}
