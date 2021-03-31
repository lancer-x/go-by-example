package main

import "fmt"

func main() {
    t1()
}

//初始化
func t1() {
    // 初始make化
	m1 := make(map[string]string)

	fmt.Println(m1)

    //字面量初始化
	m2 := map[string]string{"a": "aaa", "b": "bbb"}
    fmt.Println(m2)

    //无法用new初始化   如果错误的使用 new() 分配了一个引用对象，会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址
    m4 := new(map[string]string)
    //(*m4)["aaa"] = "bbb"
    fmt.Printf("%p, %T %T %t\n", m4, m4, *m4, *m4 == nil)
    

    //仅声明会报错
    var m3 map[string]string
    //m3["aaa"] = "bbb"
    fmt.Println(m3)
}