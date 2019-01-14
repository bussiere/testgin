package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status string `json:"Status"`
}

type Error struct {
	Message string `json:"Message"`
}

func Calculate(c *gin.Context) {
	var response Response
	fmt.Println(c.Request.URL.Query())
	fmt.Println(c.Param("foo"))
	fmt.Println(c.Param("bar"))
	response.Status = "ok"
	if c.BindJSON(&response) == nil {

		c.JSON(200, response)
	} else {
		var errorInputData Error
		errorInputData.Message = "Incorrect input data"
		c.JSON(500, errorInputData)
	}

}
func main() {
	r := gin.Default()

	r.GET("/Calculate/:foo/:bar/", Calculate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
