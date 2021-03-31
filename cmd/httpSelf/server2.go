package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data string `json:"data,omitempty"`
}

func main() {
    http.HandleFunc("/aaa", func(w http.ResponseWriter, r *http.Request) {
        data, _ := ioutil.ReadAll(r.Body)
        
        ret := &result{
            Code: 0,
            Msg: "success",
            Data: string(data),
        }

        a,err := json.Marshal(ret)
        if err != nil {
            fmt.Println(err.Error())
        }
        fmt.Println(string(a))
        w.Write(a)
    })
    http.ListenAndServe(":9999", nil)
    println("server started")
}