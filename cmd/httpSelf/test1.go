package main

import "net/http"

func main()  {
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":11111", nil)
}



func sayHello(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("lixiaoru"))
}
