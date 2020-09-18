package main

import "testing"

func TestHello(t *testing.T)  {
	got := Hello("harry")
	want := "hello,harry"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
