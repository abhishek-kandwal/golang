package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"goConnect/cmd/database"
)

func getdata(c *gin.Context) {

	dbConn := database.GetConn()

	results, err := database.DbGet(dbConn)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(results)

	c.JSON(200, gin.H{
		"message": "data",
		"lat": results.Lat,
		"lon": results.Lon,
		"country": results.Country,
	})
}
