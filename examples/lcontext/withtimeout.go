package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r *http.Response
	err error
}

func process()  {
	ctx,cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	tr := &http.Transport{}
	client := &http.Client{Transport:tr}

	c := make(chan Result, 1)
	req,err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		return
	}

	go func() {
		time.Sleep(time.Second)
		resp, err := client.Do(req)
		pack := Result{resp, err}
		c <- pack
	}()

	select {
	case <- ctx.Done():

		fmt.Println("timeout")
		case res := <-c:
		defer res.r.Body.Close()
		out,_ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("response", out)


	}
	return
}

func main8()  {
	process()
}


