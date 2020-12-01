package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main()  {
	mediaFile := os.Args[1]
	columnNum,_ := strconv.Atoi(os.Args[2])
	topFile := os.Args[3]

	mediaMap := readMediaFile(mediaFile, columnNum)
	topMap := readTopFile(topFile)


	a,b := showDiff(mediaMap, topMap)

	fmt.Println(len(mediaMap), len(topMap))
	fmt.Println(a)
	fmt.Println(b)
}

func showDiff(map1,map2 map[string]struct{}) ([]string,[]string) {
	diff1 := make([]string, 0)
	diff2 := make([]string, 0)

	for k,_ := range map1 {
		if _,ok := map2[k]; !ok {
			diff1 = append(diff1, k)
		}
	}

	for k,_ := range map2 {
		if _,ok := map1[k]; !ok {
			diff2 = append(diff2, k)
		}
	}
	return diff1, diff2
}

func readMediaFile(fileName string, columnNum int) map[string]struct{} {
	data := readFile(fileName)
	ret := make([]string, 0)

	for _,v := range data {
		columns := strings.Split(v, "\t")
		if len(columns) <= columnNum {
			continue
		}
		ret = append(ret, columns[columnNum])
	}

	return slice2map(ret)
}

func readTopFile(file string) map[string]struct{}  {
	ret := readFile(file)
	return slice2map(ret)
}


func readFile(filePath string) []string {
	data := make([]string, 0)
	fd, _ := os.Open(filePath)
	defer fd.Close()
	reader := bufio.NewReader(fd)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		data = append(data, string(line))
	}
	return data
}

func slice2map(data []string) map[string]struct{} {
	ret := make(map[string]struct{})
	for _,v := range data {
		ret[v] = struct {}{}
	}
	return ret
}
