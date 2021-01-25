package main

import "fmt"

type errNotFound1 struct {
	file string
}

func (e errNotFound1) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

func open(file string) error {
	return errNotFound{file: file}
}

func use() {
	if err := open("aa"); err != nil {
		if _, ok := err.(errNotFound); ok {
			// handle
		} else {
			panic("unknown error")
		}
	}
}