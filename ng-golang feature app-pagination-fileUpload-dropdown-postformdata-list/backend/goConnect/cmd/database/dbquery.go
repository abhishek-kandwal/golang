package database

import (
	"fmt"
	"goConnect/cmd/models"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var totalrows int64 = 5

func DbGet(db *sql.DB, action string, startIn int64 , stopIn int64) ([]models.Weather, error , int64 , int64 , int) {
	fmt.Println(action)
	fmt.Println("getting data from the db")
	
	var userData []models.Weather

	rows ,  err := db.Query("select city, conds , humidity , temp from weather ")

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var userDatas models.Weather

	for rows.Next() {
		var city string
		var conds string
		var humidity int
		var temp int
        err := rows.Scan(&city , &conds ,&humidity , &temp )
        if err != nil {
            fmt.Println(err)
		}
		userDatas.City = city
		userDatas.Conds = conds
		userDatas.Humidity = humidity
		userDatas.Temp = temp
        userData = append(userData, userDatas)
	}
	totallen := len(userData)
	if stopIn > int64(totallen) { 
		userData = userData[startIn:totallen]
	} else{
		userData = userData[startIn:stopIn]
	} 
	
	return userData, err , startIn , stopIn , totallen
}
