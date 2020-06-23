package database

import (
	"fmt"
	"goConnect/cmd/models"
	"time"
)

func GetUserDetails(email string) ( []models.UserDetail , int ,  error){
	db := GetConn();

	var userDetail []models.UserDetail
	
	var dbQueryString string

	if email != "" {
		dbQueryString ="select * from usersDetail where email=" +   email 
	}else {
		dbQueryString = "select * from usersDetail"
	}
	
	rows , err := db.Query(dbQueryString)

	if err != nil {
		panic( err.Error())
	}

	defer rows.Close()

	var user models.UserDetail

	for rows.Next(){
		var id int
		var name string
		var email string
		var dob string
		err := rows.Scan(&id, &name , &email , &dob)
		if err != nil {
			panic(err.Error())
		}
		user.ID  = id
		user.Name = name
		user.Email = email
		user.Dob = dob

		userDetail = append(userDetail , user)
	}
	totalrecords := len(userDetail)

	return userDetail, totalrecords, err
}


func SetUserDetail(name string, email string, dob string) (string , error) {
	
	db := GetConn();

	var err error
	var userDetail []models.UserDetail

	// checks 
	if name == "" {
		return "name field is required" , err
	}else if email == "" {
		return "email field is required" , err 
	}else if dob == "" {
		return "dob field is required" , err 
	}

	rows, err := db.Query("select * from usersDetail where email= \" " + email + " \" " )

	if err != nil {
		panic(err.Error())
	}

	var user models.UserDetail

	for rows.Next(){
		var id int
		var name string
		var email string
		var dob string
		err := rows.Scan(&id, &name , &email , &dob)
		if err != nil {
			panic(err.Error())
		}
		user.ID  = id
		user.Name = name
		user.Email = email
		user.Dob = dob

		userDetail = append(userDetail , user)
	}
	totalrecords := len(userDetail)

	if totalrecords > 0 {
		fmt.Println("already")
	}

	date , _ := time.Parse("2006-01-02", dob[0:10] ) // slicing the string to get the format of "2020-06-08" from "2020-06-08T18:30:00.000Z".

	insert , err := db.Query("insert into usersDetail (name , email , dob) values (?,?,?)", name , email , date )

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	return "Success" , err
	
}