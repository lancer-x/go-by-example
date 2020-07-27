package main

import (
	"fmt"
	"strconv"
)

func main()  {
	a := lprint("aaaa")
	fmt.Println(a)

	fmt.Println(lprint(1111))

}

func lprint(value interface{}) string {
	type stringer interface {
		String() string
	}

	switch t := value.(type) {
	case stringer:
		return t.String()
	case string:
		return t
	case int:
		return strconv.Itoa(t)
	case bool:
		if t {
			return "true"
		} else {
			return "false"
		}
	default:
		return "???"
	}
}