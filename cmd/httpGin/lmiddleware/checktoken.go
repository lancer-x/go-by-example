package lmiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TokenChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("token checked")
	}
}
