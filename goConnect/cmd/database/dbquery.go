package database

import (
	"fmt"
	"goConnect/cmd/models"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbInsert(db *sql.DB, dataValue models.Weather) bool {

	fmt.Println("Inserting Data")

	insert, err := db.Query("INSERT INTO weathermap (name, visibility, base, dt,cond,lon,lat) VALUES (?,?,?,?,?,?,?)", dataValue.Name, dataValue.Visibility, dataValue.Base, dataValue.Dt, dataValue.Country, dataValue.Lon, dataValue.Lat)

	if err != nil {
		panic(err.Error())
		return false
	}

	defer insert.Close()

	fmt.Println("Succefully Inserted!")
	return true
}

func DbGet(db *sql.DB) (models.Weather, error) {
	fmt.Println("getting data from the db")
	var userData models.Weather
	err := db.QueryRow("select id ,name ,visibility ,base ,dt ,cond ,lon ,lat from weathermap WHERE cond='GB'").Scan(&userData.Id3, &userData.Name, &userData.Visibility, &userData.Base, &userData.Dt, &userData.Country, &userData.Lon, &userData.Lat)

	if err != nil {
		panic(err.Error())
	}

	// for result.Next() {

	// 	err = result.Scan(&userData.Name, &userData.Id1)

	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// fmt.Printf("name: %s\n", userData.Name)

	// return userData, err
	//}
	return userData, err
}
