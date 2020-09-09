package main

import (
	"encoding/json"
	"go-by-example/cmd/httpGin/lmiddleware"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.Use(lmiddleware.TokenChecker())
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message" : "pong",
			})
		})
	}



	r.POST("/colorjson", func(c *gin.Context) {
		c.JSON(200, gin.H{"str": "str-val", "int": 1})
	})
	r.GET("/testget", func(c *gin.Context) {
		c.JSON(200, gin.H{"method": "get", "code": 0})
	})

	r.Any("/backraw", func(c *gin.Context) {
		bodyRaw,_ := c.GetRawData()
		var body map[string]interface{}
		if err := json.Unmarshal(bodyRaw, &body); err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
		}
		header := c.Request.Header

		c.JSON(200, gin.H{"url": c.Request.URL.String(), "header": header, "body": body})
	})
	r.Run("127.0.0.1:11111")
}