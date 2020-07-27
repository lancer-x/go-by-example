package main

import "net/http"

func main()  {
	http.HandleFunc("/info1", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("info1"))
	})

	http.ListenAndServe(":11113", nil)
}