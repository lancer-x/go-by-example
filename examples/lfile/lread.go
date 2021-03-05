package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main3()  {
	file := "/Users/lightli/data/userphone.csv"
	c1 := read1(file)
	c2 := read2(file)
	c3 := read3(file)
	fmt.Println(c1, c2, c3)
}

func read1(filename string) int {
	ff, _ := os.Open(filename)

	defer ff.Close()
	buf := make([]byte, 4096)
	var byteCount  int

	for {
		n, _ := ff.Read(buf)
		if n == 0 {
			break
		}
		byteCount += n
	}
	return byteCount
}

func read2(filename string) int {
	ff, _ := os.Open(filename)
	defer ff.Close()
	buf := make([]byte, 4096)
	var byteCount int

	rd := bufio.NewReader(ff)

	for {
		n,_ := rd.Read(buf)
		if n == 0 {
			break
		}
		byteCount += n
	}
	return byteCount
}

func read3(filename string) int {
	ff,_ := os.Open(filename)
	defer ff.Close()

	fd,_ := ioutil.ReadAll(ff)
	return len(fd)
}

