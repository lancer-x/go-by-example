package main

import (
	"encoding/json"
	"fmt"
)

type Transport struct {
	Time string
	Mac string
	Id string
	Rssid string
}

func main()  {
	var st []Transport
	t1 := Transport{
		Time:  "11",
		Mac:   "22",
		Id:    "33",
		Rssid: "44",
	}
	t2 := Transport{
		Time:  "55",
		Mac:   "66",
		Id:    "77",
		Rssid: "88",
	}

	st = append(st, t1, t2)
	buf, _ := json.Marshal(st)
	fmt.Println(string(buf))

	var st1 []Transport
	json.Unmarshal(buf, &st1)
	fmt.Println(st1)

	var st2 []map[string]string
	json.Unmarshal(buf, &st2)
	fmt.Println(st2)
}