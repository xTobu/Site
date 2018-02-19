package handlersApi

import (
	"net/http"

	// "./mssqldb"

	"./mysqldb"
	"github.com/gin-gonic/gin"
)

//StudentStruct struct
type structStudent struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	CreatedTime string `json:"CreatedTime,omitempty"`
}

//Student3 use mysql
func Student3(c *gin.Context) {
	if res, count := mysqldb.DBGetStudents(); count > 0 {
		n := mysqldb.StudentS{res}
		c.JSON(http.StatusOK, n)
	} else {
		n := mysqldb.StudentS{}
		c.JSON(http.StatusOK, n)
		// c.JSON(http.StatusNoContent)
	}
	// if res, count := sqldb.DBGetStudents(); res != nil {
	// 	if cou := count; cou >= 0 {
	// 		n := sqldb.StudentS{res}
	// 		c.JSON(http.StatusOK, n)
	// 	}
	// }
}
