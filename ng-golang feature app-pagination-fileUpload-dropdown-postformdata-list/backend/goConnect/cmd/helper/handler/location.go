package handler

import (
	"goConnect/cmd/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

type dataStruct struct {
	name string
	id int
}

func getlocation(c *gin.Context){

	dbConn := database.GetConn()
	
	state, _ := strconv.ParseInt(c.Query("state"), 10, 64)

	results , err := database.GetLocation(dbConn , state)

	if err != nil {
		panic(err.Error())
	}
	
	c.JSON(200, gin.H{
		"data": results,
	})
}