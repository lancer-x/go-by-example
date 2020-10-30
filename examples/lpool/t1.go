package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type person struct{
	name string
	age int
}

const loopNum  = 100000000


func main()  {
	fmt.Println(strconv.ParseFloat(fmt.Sprintf("%.2f", 1.467), 64))
	//testCommon()
	//testPool()
}

func testCommon()  {
	begin := time.Now()
	var obj *person
	for i := 0; i<loopNum; i++ {
		obj = &person{
			name: "test",
			age:  0,
		}
		obj.age++
	}
	end := time.Since(begin)
	println(int(end))
}

func testPool()  {
	intPool := sync.Pool{
		New: func() interface{} {
			return &person{
				name: "base",
				age:  10,
			}
		},
	}

	var obj *person
	begin := time.Now()
	for i := 0; i<loopNum; i++ {
		obj = intPool.Get().(*person)
		obj.age = i
		intPool.Put(obj)
	}
	end := time.Since(begin)
	println(int(end))
	println()

}


