package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const turl = "http://127.0.0.1:9999/test?a=1&b=2&c=3"

func main() {
	t4()
}

func t1() {
	resp, _ := http.Get(turl)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func t2() {
	req, _ := http.NewRequest("GET", turl, nil)

	req.Header.Add("aaa", "bbb")
	http.DefaultClient.Do(req)
}

type t2body struct {
	Name   string
	Age    int
	Salary int
}

func t3() {
	body := &t2body{
		Name:   "light",
		Age:    20,
		Salary: 100,
	}
	data, _ := json.Marshal(*body)
	req, _ := http.NewRequest("POST", turl, bytes.NewBuffer(data))
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	req.Header.Add("If-AAA", `W/vvv"`)
	http.DefaultClient.Do(req)
}

func t4() {
	http.PostForm(turl, url.Values{"key": {"Value"}, "id": {"123"}})
}
