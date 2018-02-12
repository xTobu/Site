package handlersApi

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"./sqldb"
	"github.com/gin-gonic/gin"
)

// //StudentStruct struct
// type structStudent struct {
// 	ID          int    `json:"id,omitempty"`
// 	Name        string `json:"name,omitempty"`
// 	Email       string `json:"email,omitempty"`
// 	CreatedTime string `json:"CreatedTime,omitempty"`
// }

//Insert handler 新增一筆學生資料
func Insert(c *gin.Context) {
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")

	//開啟db
	db, err := sql.Open("mssql", sqldb.ConnectionStr)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	query := fmt.Sprintf("INSERT INTO [dbo].[vueStudent] ([name],[email]) VALUES('%s','%s')", name, email)
	result, err := db.Query(query)
	defer result.Close()

	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}

//Insert2 handler 新增一筆學生資料
func Insert2(c *gin.Context) {

	if r := sqldb.DBInsertStudent(c.Request.FormValue("name"), c.Request.FormValue("email")); r == true {
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "fail",
		})
		// c.JSON(http.StatusNoContent)
	}
}
