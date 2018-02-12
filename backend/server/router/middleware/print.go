package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// PrintMiddleware is a function for test middleware
func PrintMiddleware(c *gin.Context) {
	fmt.Println("PrintMiddleware")
	c.Next()
}

// Print is a function for test middleware
func Print() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Print")
		c.Next()
	}
}
