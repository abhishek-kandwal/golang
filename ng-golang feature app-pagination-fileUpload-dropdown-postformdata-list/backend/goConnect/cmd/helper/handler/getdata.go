package handler

import (
	// "fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"goConnect/cmd/database"
)

func getdata(c *gin.Context) {

	dbConn := database.GetConn()

	action := c.Query("action")
	// start := c.Query("startIndex")
	// stop := c.Query("stopIndex")

	start , _ := strconv.ParseInt(c.Query("startIndex"), 10, 64)
	stop , _ := strconv.ParseInt(c.Query("stopIndex"), 10, 64)

	results, err, startat , stopat , length := database.DbGet(dbConn, action, start, stop)

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, gin.H{
		"message": "data",
		"data":    results,
		"start":    startat,
		"stop":    stopat,
		"total": length,
	})
}
