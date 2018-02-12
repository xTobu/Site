package sqldb

import (
	"database/sql"
	"fmt"

	"log"
)

func dbGetConn() *sql.DB {
	db, err := sql.Open("mssql", ConnectionStr)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}

	return db
}

// DBGetStudents 取得所有學生資料
func DBGetStudents() (data []Student, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("select * from [vueStudent]")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var student Student
		rows.Scan(&student.ID, &student.Name, &student.Email, &student.CreatedTime)
		data = append(data, student)
		count++
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
	getTime()
	return data, count
}

// DBInsertStudent 新增一筆學生資料
func DBInsertStudent(name string, email string) (r bool) {

	// name := c.Request.FormValue("name")
	// email := c.Request.FormValue("email")
	//name = "A"
	//email = "B"

	//開啟db
	db := dbGetConn()
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO [dbo].[vueStudent] ([name],[email],[CreatedTime]) VALUES('%s','%s','%s')", name, email, getTime())
	result, err := db.Query(query)
	defer result.Close()

	if err != nil {
		log.Fatalln(err)
	}

	r = true
	return r

}
