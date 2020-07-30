package main

import (
	"context"
	"fmt"
)

func main()  {
	test1()
}

func test1()  {
	ctx := context.Background()

	ctx1 := context.WithValue(ctx, "name", "张三")

	fmt.Println(ctx1.Value("name"))
}