package handler

import (
	"fmt"
	// "io/ioutil"
	"github.com/gin-gonic/gin"
)

func Handler() {

	fmt.Println("in Handler")
	r := gin.Default()

	r.GET("/", homepage)
	r.GET("/getdata", getdata)
	r.GET("/fetchapi", fetchapi)
	
	r.Run()
}

func homepage(c *gin.Context){
	c.JSON(200, gin.H{
		"message":"welcome to goConnect",
	})
}