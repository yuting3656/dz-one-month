package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Wrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("************************************")
		c.Next()
		fmt.Println("************************************")
	}
}
