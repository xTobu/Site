package handlersVue

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Student handler
func Student(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{
		"student": "Student",
	})
}
