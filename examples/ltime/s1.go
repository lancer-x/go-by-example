package main

import (
	"fmt"
	"time"
)

func main()  {
	t1()
}

func t1() {
	a := time.Now()
	fmt.Println(a)
	fmt.Println(a.Unix())

	fmt.Println(a.Format("2006-01-02 15:04:05"))


	fmt.Println(a.Truncate(1 * time.Minute), a.Truncate(1 * time.Hour))

	b := time.Unix(1611372296, 0)
	fmt.Println(b)

	fmt.Println(a.After(b))
}


