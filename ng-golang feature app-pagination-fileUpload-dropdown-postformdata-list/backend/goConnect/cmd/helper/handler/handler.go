package handler

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Handler() {

	fmt.Println("in Handler")
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // http://localhost:4200
		AllowMethods:     []string{"GET", "POST",  "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin","Content-Type","application/json"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*" // http://localhost:4200
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/", homepage)
	r.GET("/weather", getdata)
	r.GET("/location", getlocation)
	r.GET("/employee", getEmployee)
	r.GET("/guserdetails", getUserDetails)
	
	
	r.POST("/puserdetails", postUserDetails)
	r.POST("/fileUpload", fileUpload)

	r.Run()
}

func homepage(c *gin.Context){
	c.JSON(200, gin.H{
		"message":"welcome to goConnect",
	})
}