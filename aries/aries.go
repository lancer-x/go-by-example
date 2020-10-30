package aries

import (
	"log"
	"net/http"
)

//全局唯一请求处理
type AriesContext struct {
	handlerMap map[string]http.HandlerFunc
	uri404Handler http.HandlerFunc
}

func NewAries() *AriesContext {
	return &AriesContext{
		handlerMap:make(map[string]http.HandlerFunc),
	}
}

func (ac *AriesContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}

	uri := r.URL.Path
	log.Println("calling", uri)
	if handler,ok := ac.handlerMap[uri]; ok {
		handler(w, r)
	} else {
		if ac.uri404Handler != nil {
			ac.uri404Handler(w, r)
		} else {
			_, _ = w.Write([]byte("uri not found"))
		}
	}

}

func (ac *AriesContext) Get(uri string, h http.HandlerFunc) {
	ac.handlerMap[uri] = h
}

func (ac *AriesContext) Uri404 (h http.HandlerFunc)  {
	ac.uri404Handler = h
}

func (ac *AriesContext) Run(addr string) error {
	return http.ListenAndServe(addr, ac)
}