package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"goConnect/cmd/database"
	"goConnect/cmd/helper/convert"
	"goConnect/cmd/helper/mapping"

	"github.com/gin-gonic/gin"
)

func fetchapi(c *gin.Context) {
	fmt.Println("fetching data")

	response, err := http.Get("https://samples.openweathermap.org/data/2.5/weather?q=London,uk&appid=439d4b804bc8187953eb36d2a8c26a02")

	if err != nil {
		panic(err.Error)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err.Error())
	}

	result, err := convert.ByteToInterface(body)

	if err != nil {
		panic(err.Error())
	}

	DataModel, err := mapping.DataMap(result)

	if err != nil {
		panic(err.Error())
	}

	dbConn := database.GetConn()

	if dbConn == nil {
		panic("no connection to db")
	}

	isInserted := database.DbInsert(dbConn, DataModel)

	if isInserted {
		c.JSON(200, gin.H{
			"message": "data Inserted",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "data Not Inserted",
		})
	}
}
