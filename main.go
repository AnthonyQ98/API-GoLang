package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type phone struct{
	ID       string `json:"id"`
	Model    string `json:"model"`
	Year     string `json:"year"`
	Quantity int    `json:"quantity"`
}

var phones = []phone{
	{ID: "1", Model: "iPhone 11", Year: "2019", Quantity: 4},
	{ID: "2", Model: "iPhone 6", Year: "2014", Quantity: 9},
	{ID: "3", Model: "iPhone X", Year: "2017", Quantity: 2},
}

func getPhones(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, phones)
}

func createPhone(c *gin.Context) {
	var newPhone phone

	if err := c.BindJSON(&newPhone); err != nil {
		return 
	}

	phones = append(phones, newPhone)
	c.IndentedJSON(http.StatusCreated, newPhone)
}

func main(){
	router := gin.Default()
	router.GET("/phones", getPhones)
	router.POST("/phones", createPhone)
	router.Run("localhost:8080")
}