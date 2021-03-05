package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

type httpData struct {
	Body map[string]interface{} `json:"body"`
	Header map[string]interface{} `json:"header"`
	Url string `json:"url"`
}

type user struct {
	Name string
	Age int
}

var json1 = `{
    "body": {
        "data": {
            "aaa_cc": "aaa",
            "bbb": "bbbb"
        },
        "str": "fooo"
    },
    "header": {
        "Accept-Encoding": [
            "gzip"
        ],
        "Content-Length": [
            "48"
        ],
        "Content-Type": [
            "application/json"
        ],
        "Header-A": [
            "header-a"
        ],
        "Header-B": [
            "2"
        ],
        "User-Agent": [
            "Go-http-client/1.1"
        ]
    },
    "url": "/backraw?q1=1\u0026q2=qq2"
}`

func maine()  {
	map1 := make(map[string]interface{})
	var map2 httpData

	//json  to map
	json.Unmarshal([]byte(json1), &map1)

	//map to struct
	mapstructure.Decode(map1, &map2)

	fmt.Println(map1)
	fmt.Println(map2)

	user1 := user{
		Name: "郭德纲",
		Age:  10,
	}

	user1Map := struct2Map(user1)
	fmt.Println(user1Map)
}

func struct2Map(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)

	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

