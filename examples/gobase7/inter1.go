package main

import (
	"fmt"
	"io"
)

type demoWriter struct {
	buf []byte
}

func (d demoWriter) Write(con []byte) (int,error) {
	d.buf = append(d.buf, con...)
	len := len(d.buf)
	return len, nil
}

func (d *demoWriter) Read(con []byte) (int, error) {
	return 0, nil
}

func (d demoWriter) showLen() {
	fmt.Printf("%d---\n", len(d.buf))
}

func main1()  {
	var w io.Writer
	var r io.Reader
	d := &demoWriter{}
	w = d

	r = &demoWriter{}

	len,_ := w.Write([]byte{'1','2','3'})
	fmt.Println(len, r)

	d.showLen()

}


