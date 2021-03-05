package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var result = make(map[string]int)

func main1()  {
	runMore()
	fmt.Println(result)
}

func runSync()  {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func runMore()  {
	wg := sync.WaitGroup{}
	count := len(os.Args[1:])
	wg.Add(count)
	for _, url := range os.Args[1:]{
		go func(url string) {
			s := getUrl(url)
			setMap(url, &s)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func getUrl(url string)(s []byte) {
	resp,_ := http.Get(url)
	s,_ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return
}

func setMap(url string, s *[]byte)  {
	ml := sync.Mutex{}
	ml.Lock()
	{
		if _,ok := result[url]; ok {
			url += strconv.Itoa(rand.Int())
			fmt.Println(url)
		}
		result[url] = len(*s)
	}
	ml.Unlock()
}
