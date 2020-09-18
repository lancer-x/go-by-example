package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main()  {
	
	handle()
}

func handle()  {
	v,_ := strconv.Atoi(os.Args[1])
	fileName := "./" + os.Args[2] + ".log"

	switch v {
	case 1:
		write1(fileName)
	case 2:
		write2(fileName)
	case 3:
		write3(fileName)
	case 4:
		write4(fileName)
	}
}

func write1(fileName string)  {
	var fd *os.File
	if checkFileExist(fileName) {
		fd,_ = os.OpenFile(fileName, os.O_APPEND, 0666)

	} else {
		fd,_ = os.Create(fileName)
	}
	fd.Close()

	n,_ := io.WriteString(fd, "asfjakfj")
	fmt.Println(n)
}

func write2(fileName string)  {
	bytes := []byte("afjakjlg")
	_ =ioutil.WriteFile(fileName, bytes, 0666)
}

func write3(fileName string) {
	var fd *os.File
	if checkFileExist(fileName) {
		fd,_ = os.OpenFile(fileName, os.O_APPEND, 0666)

	} else {
		fd,_ = os.Create(fileName)
	}
	defer fd.Close()
	_, _ = fd.WriteString("aaaaaa")
}

func write4(fileName string)  {
	var fd *os.File
	if checkFileExist(fileName) {
		fd,_ = os.OpenFile(fileName, os.O_APPEND, 0666)

	} else {
		fd,_ = os.Create(fileName)
	}
	defer fd.Close()

	w := bufio.NewWriter(fd)
	_, _ = w.WriteString("asfjaljasglkds")
	_ = w.Flush()
}

func checkFileExist(fileName string) bool {
	var exist = true
	if _,err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}