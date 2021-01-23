package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main()  {
	//t1()
	t3()
}

func t0()  {
	b := []byte("瓜子瓜子")
	fmt.Println(b, string(b))

	//byte 转 rune
	c := bytes.Runes(b)
	fmt.Println(c, string(c))
}

func t1()  {
	b := []byte("瓜子瓜子")

	r1 := bytes.NewReader(b)
	d1 := make([]byte, len(b))
	n,_ := r1.Read(d1)
	fmt.Println(n, string(d1))

	r2 := bytes.Reader{}
	r2.Reset(b)
	d2 := make([]byte, len(b))
	n,_ = r2.Read(d2)
	fmt.Println(n, string(d2))
}

func t3()  {
	a,_ := strconv.Atoi("ab")
	fmt.Println(a)

	b := strconv.Itoa(12)
	fmt.Println(b)

}
