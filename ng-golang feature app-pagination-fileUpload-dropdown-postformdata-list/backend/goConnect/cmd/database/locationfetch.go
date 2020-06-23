package database

import (
	"database/sql"
	"fmt"
	"goConnect/cmd/models"

)

func GetLocation(db *sql.DB , state int64) ([]models.Datalocation , error){
	
	var locationData []models.Datalocation
	var dbString string

	if state > 0 {
		Query := "select id, name from cities where stateid="
		dbString = fmt.Sprintf("%s%d", Query, state)
	}else {
		dbString = "select id, name from states "
	}
	
	rows ,  err := db.Query(dbString)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var userDatas models.Datalocation

	for rows.Next() {
		var name string
		var id int64
        err := rows.Scan(&id , &name )
        if err != nil {
            fmt.Println(err)
		}

		userDatas.Id = id
		userDatas.Name = name

        locationData = append(locationData, userDatas)
	}
	
	return locationData, err 
}