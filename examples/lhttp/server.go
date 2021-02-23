package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 */
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Header", r.Header)
	fmt.Println("Query", r.URL.Query())
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Body", string(body))
	fmt.Println("Form", r.PostForm)
}

func main() {
	http.HandleFunc("/test", Test)
	http.ListenAndServe(":9999", nil)
}
