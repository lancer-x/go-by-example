package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.14
	t := reflect.TypeOf(x)

	fmt.Printf("%T%v",t, t)

	v := reflect.ValueOf(x)
	fmt.Println(v)

	aa := v.Interface().(float64)
	fmt.Println(aa)
}
