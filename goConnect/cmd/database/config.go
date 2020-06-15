package database

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	
)

var databaseConnection *sql.DB

func Conn() ( *sql.DB , error) {
	
	fmt.Println("connecting to DB")
	
	var err error

	databaseConnection, err = sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test_demo")
	
	fmt.Println("connected to db")

	return databaseConnection , err 
}

func GetConn() (*sql.DB){
	return databaseConnection
}