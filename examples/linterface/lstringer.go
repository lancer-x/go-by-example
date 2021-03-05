package main

import "fmt"

type country struct {
	name string
}

type city struct {
	name string
}

type lstringer interface {
	toString() string
}

func (c country) toString () string {
	return "country:" + c.name
}

func (c city) toString () string {
	return "city:" + c.name
}

func lprint(s lstringer)  {
	fmt.Println(s.toString())
}

func main6()  {
	aCountry := country{
		name: "中国",
	}
	aCity := city{name: "北京"}

	lprint(aCountry)
	lprint(aCity)
}