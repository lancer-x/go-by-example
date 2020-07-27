package main

import "net/http"

type Handle struct {

}

func (h Handle) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/info":
		response.Write([]byte("info"))
	default:

	}
}

func main()  {
	http.ListenAndServe("127.0.0.1:11112", Handle{})
}
