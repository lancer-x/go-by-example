package main

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/gin-gonic/gin"
	//"time"
)

const url  = "127.0.0.1:11111"
func main()  {

	//time.Sleep(time.Millisecond * 200)
	//mapExample()
	testGet()
}

func mapExample() {
	err := gout.POST(url + "/colorjson").
		Debug(true).
		SetJSON(gout.H{"str":"foo","num":100,"bool": false,}).
		Do()
	if (err != nil) {
		fmt.Printf("err = %v\n", err)
	}
}

func testGet()  {
	err := gout.GET(url + "/backraw").Debug(true).
		SetJSON(gout.H{"str": "fooo", "data" : gout.H{"aaa":"aaa", "bbb":"bbbb"}}).
		SetHeader(gout.H{"header-a":"header-a", "header-b": 2}).
		SetQuery(gout.H{"q1":1,"q2":"qq2"}).Do()

	if (err != nil) {

	}
}

func server() {
	router := gin.Default()
	router.POST("/colorjson", func(c *gin.Context) {
		c.JSON(200, gin.H{"str": "str-val", "int": 1})
	})
	router.GET("/testget", func(c *gin.Context) {
		c.JSON(200, gin.H{"method": "get", "code": 0})
	})

	router.Run(url)
}
