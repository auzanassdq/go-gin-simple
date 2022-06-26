package main

import (
	"fmt"
	"gogin/models"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello gin",
	})
}

func PostData(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Println("something error")
	}

	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func main() {
	fmt.Println("hello")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostData)
	r.GET("/movie", models.AllMovie)

	r.Run()
}
