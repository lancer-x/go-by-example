package main

import (
	"got/aries"
	"net/http"
)

func main()  {
	ac := aries.NewAries()
	ac.Uri404(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("当前uri无法处理，请重试"))
	})
	ac.Get("/hello", sayHi)
	ac.Run(":12222")
}

func init()  {

}


func sayHi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}