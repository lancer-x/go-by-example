package main

import (
	"encoding/json"
	"fmt"
)

type Phone struct {
	Ascreen Screen
}

type SecondPhone struct {
	Screen
}

type Screen struct {
	Width int
	Height int
}

func main2()  {
	phone := Phone{
		Ascreen: Screen{
			Width:  50,
			Height: 50,
		},
	}

	secondPhone := SecondPhone{Screen{
		Width:  100,
		Height: 100,
	}}
	fmt.Println(phone)
	fmt.Println(secondPhone)

	fmt.Println(phone.Ascreen.Width)
	fmt.Println(secondPhone.Width)

	data,_ := json.MarshalIndent(phone,"", "    ")
	fmt.Println(string(data))
	var s Phone
	json.Unmarshal(data, &s)
	fmt.Println(s)

}
