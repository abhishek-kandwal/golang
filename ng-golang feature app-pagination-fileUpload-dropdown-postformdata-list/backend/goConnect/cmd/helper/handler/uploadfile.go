package handler

import (
	"os"

	"github.com/gin-gonic/gin"
)

func fileUpload(c *gin.Context) {
	var message string

	id := c.Query("id")
	file, err := c.FormFile("file")

	if err != nil {
		panic(err.Error())
	}

	dir, err := os.Getwd()
	
	if err != nil {
		panic(err.Error())
	}

	//  checking is file exist or not
	_ , err = os.Stat( dir + "/FileStorage/" + id + "/" + file.Filename)

    if os.IsNotExist(err) {
		// if this file not exist 

		os.Mkdir("FileStorage/" + id , 0755)
		
		dir = dir + "/FileStorage/" + id + "/"
		path := dir + file.Filename
		
		if err := c.SaveUploadedFile(file, path); err != nil {
			message = "some error occured"
			panic(err.Error())
		}else {
			message = "Successfully Uploaded"
		}
		
		} else {
			// if already exist
			message = "File Already exist"
	}


	c.JSON(200, gin.H{
		"isUpload": message,
	})
}
