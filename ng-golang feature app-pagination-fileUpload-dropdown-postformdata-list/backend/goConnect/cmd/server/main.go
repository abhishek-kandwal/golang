package main

import (
	"fmt"

	"goConnect/cmd/database"
	"goConnect/cmd/helper/handler"
)

// type response1 struct{}

func main(){ 
	db , err := database.Conn()
	
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("connected to db")
	
	defer db.Close()

	handler.Handler()

}