package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Hello handler
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}

//Get handler
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get",
	})
}
