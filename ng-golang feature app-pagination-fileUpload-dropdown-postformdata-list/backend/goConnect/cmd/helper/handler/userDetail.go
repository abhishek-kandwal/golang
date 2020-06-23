package handler

import (
	"fmt"
	"io/ioutil"
	"goConnect/cmd/helper/convert"

	"github.com/gin-gonic/gin"

	"goConnect/cmd/database"
)

func getUserDetails(c *gin.Context){
	email := c.Query("email");
	
	results, datalen , err := database.GetUserDetails(email);

	if err != nil {
		panic(err.Error())
	}

	if email != ""{
		isExist := false
		
		if datalen > 0 {
			isExist = true
		}

 	    c.JSON(200, gin.H{
			"isExist": isExist,
		})
	}else {
		c.JSON(200, gin.H{
			"userdata": results,
		})
	}

}

type userstruct struct {
		name string `json:"name" binding:"required"`
		email string `json:"email" binding:"required"`
		dob string `json:"dob" binding:"required"`
}

func postUserDetails(c *gin.Context){

	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	data , err := convert.ByteToInterface(x)
	
	if err != nil {
		panic(err.Error())
	}

	 userdata := userstruct{
		 name : fmt.Sprintf("%v", data["name"]),
		 email: fmt.Sprintf("%v" , data["email"]),
		 dob: fmt.Sprintf("%v" , data["dob"]),
	 }
	 

	message , err := database.SetUserDetail(userdata.name , userdata.email , userdata.dob )

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200 , gin.H{
		"message": message,
	})
}