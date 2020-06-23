package database

import (
	"fmt"
	"goConnect/cmd/models"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)




func GetEmoloyeeList(db *sql.DB,) ([]models.Employee, error ) {

	var userData []models.Employee

	rows ,  err := db.Query("select id , name from employee ")

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var userDatas models.Employee

	for rows.Next() {
		var id int
		var name string
        err := rows.Scan(&id , &name )
        if err != nil {
            fmt.Println(err)
		}
		userDatas.Id = id
		userDatas.Name = name
        userData = append(userData, userDatas)
	}
	
	return userData, err
}
