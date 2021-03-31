package main

import "net/url"

func main() {
	a := "哈哈"

	b, err := url.QueryUnescape(a)
	c := url.QueryEscape(a)
    if err != nil {
        println(err.Error())
    }
    print(b)
    print(c)
}