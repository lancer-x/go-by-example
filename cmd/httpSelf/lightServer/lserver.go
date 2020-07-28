package lightServer

import "net/http"

type LResponseWriter struct {

}

func (lr LResponseWriter) Header()  *http.Header{
	return new(http.Header)
}