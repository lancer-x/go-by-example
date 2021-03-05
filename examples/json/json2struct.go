package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type params struct {
	Happy  string                 `json:"happy"`
	Id     string                 `json:"id"`
	Query  []string               `json:"query"`
	Fields map[string]interface{} `json:"fields"`
	//TmpTest  string                 `json:"tmp_test"` 错误的
	//TmpTests string                 `json:"tmpTests"` 正确的
}

func mainf() {
	tmp := make(map[string]interface{})
	//tmp["id"] = "123"
	//tmp["query"] = []string{"qwe", "wer", "ert"}

	param := `
			{
				"happy": "sds",
				"query":[
					"qwde",
					"wer"
				],
   				"fields":{
       			"name":"Doria",
       			"love":"HelloKitty"
   			},
				"id":"1234"
			}
`

	var confg params
	json.Unmarshal([]byte(param), &tmp)
	err := mapstructure.Decode(tmp, &confg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("get the map : ",tmp)
	fmt.Println("get the struct : ",confg)
}

