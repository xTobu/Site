package handlersVue

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	//"../../../mssql"
	//go mssql
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

//GetData handler
func GetData(c *gin.Context) {
	//mssql.Init()
	// c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{
		"student": "Student",
	})

}

//User struct
type User struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	CreatedTime string `json:"CreatedTime,omitempty"`
}

//GetUser func
func GetUser(w http.ResponseWriter, req *http.Request) {
	var r User
	// params := mux.Vars(req)

	db, err := sql.Open("mssql", "server=211.78.85.202;database=junxiang_db;user id=junxiang;password=WPE8yxwq;encrypt=disable;")
	if err != nil {
		log.Fatal(err)
	}

	// err1 := db.QueryRow("select Id, UserName from [Your Datavse].dbo.Users where Id = ?", params["id"]).Scan(&r.ID, &r.Name)
	err1 := db.QueryRow("select * from [junxiang_db].[dbo].[vueStudent]").Scan(&r.ID, &r.Name, &r.Email, &r.CreatedTime)
	if err1 != nil {
		log.Fatal(err1)
	}

	json.NewEncoder(w).Encode(&r)

	if err != nil {
		log.Fatal(err)
	}

}
