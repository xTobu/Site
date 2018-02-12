package handlersApi

import (
	"database/sql"
	"log"
	"net/http"

	"./sqldb"
	"github.com/gin-gonic/gin"
)

//StudentStruct struct
type structStudent struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	CreatedTime string `json:"CreatedTime,omitempty"`
}

//Student handler
func Student(c *gin.Context) {
	db, err := sql.Open("mssql", sqldb.ConnectionStr)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	rows, err := db.Query("select * from [vueStudent]")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	data := make([]structStudent, 0)

	for rows.Next() {
		var student structStudent
		rows.Scan(&student.ID, &student.Name, &student.Email, &student.CreatedTime)
		data = append(data, student)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"student": data,
	})

}

//Student2 blabla
func Student2(c *gin.Context) {
	if res, count := sqldb.DBGetStudents(); count > 0 {
		n := sqldb.StudentS{res}
		c.JSON(http.StatusOK, n)
	} else {
		n := sqldb.StudentS{}
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
