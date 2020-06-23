package handler

import (
	"goConnect/cmd/database"

	"github.com/gin-gonic/gin"
)

func getEmployee(c *gin.Context){
	dbConn := database.GetConn()

	result , err := database.GetEmoloyeeList(dbConn)

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200 , gin.H{
		"employeeList" : result,
	})
}