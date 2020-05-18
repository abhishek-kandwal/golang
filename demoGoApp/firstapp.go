package main

import (
	"io/ioutil"
	"fmt"

	"github.com/gin-gonic/gin"
)

func homepage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func QueryParam(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name":    name,
		"age":     age 
	})
}

func pathQueryParam(c *gin.Context){
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func postmethod( c *gin.Context){
	c.JSON(200 , gin.H{
		"message": "post method hit",
	})
}

func bodyResponse( c *gin.Context){
	body := c.Request.Body
	val, err := ioutil.ReadAll(body)
	if err != nil{
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"message": "bodyreader",
		"body-response": val,
	})
}

func main() {
	//  var i int=42

	fmt.Println("hello world")
	r := gin.Default()

	r.GET("/", homepage)
	r.GET("/login", login)
	r.GET("/query", QueryParam) // localhost:8080/query?name="tony"&age=26
	r.get("/pathParam/:name/:age", pathQueryParam) // localhost:8080/pathParam/tony/26
	r.POST("/", postmethod)
	r.POST("/bodyresponse", bodyResponse) // this'll response anything send to the POST request

	r.Run()
}
