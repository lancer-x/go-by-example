package main

import (
	"encoding/json"
	"fmt"
)

type CaInfo struct {
	CaS         string `json:"ca_s"`
	CaN         string `json:"ca_n"`
	CaMedium    string `json:"ca_medium"`
	CaTerm      string `json:"ca_term"`
	CaContent   string `json:"ca_content"`
	CaCampaign  string `json:"ca_campaign"`
	CaKw        string `json:"ca_kw"`
	Keyword     string `json:"keyword"`
	CaKeywordid string `json:"ca_keywordid"`
	Scode       string `json:"scode"`
	CaTransid   string `json:"ca_transid"`
	Platform    string `json:"platform"`
	Version     string `json:"version"`
	TrackId     string `json:"track_id"`
	CaCreativeid     string `json:"ca_creativeid,omitempty"`
}

func main()  {
	a := &CaInfo{
		CaS:          "11",
		CaN:          "222",
		CaMedium:     "33",
		CaTerm:       "44",
		CaContent:    "55",
		CaCampaign:   "55",
		CaKw:         "",
		Keyword:      "",
		CaKeywordid:  "",
		Scode:        "55",
		CaTransid:    "",
		Platform:     "55",
		Version:      "",
		TrackId:      "",
		CaCreativeid: "3333",
	}

	ajson, _ := json.Marshal(a)
	fmt.Println(string(ajson))
}
