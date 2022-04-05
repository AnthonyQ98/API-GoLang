package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func phoneById(c *gin.Context) {
	id := c.Param("id")
	phone, err := getPhoneById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Phone not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, phone)
}

func checkoutPhone(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Missing id query paramater"})
		return
	}

	phone, err := getPhoneById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Phone not found"})
		return
	}

	if phone.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Phone not available."})
		return
	}

	phone.Quantity -= 1
	c.IndentedJSON(http.StatusOK, phone)
}

func returnPhone(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Missing id query paramater"})
		return
	}

	phone, err := getPhoneById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Phone not found"})
		return
	}

	if phone.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Phone not available."})
		return
	}

	phone.Quantity += 1
	c.IndentedJSON(http.StatusOK, phone)
}

func getPhoneById(id string) (*phone, error) {
	for i, p := range phones {
		if p.ID == id {
			return &phones[i], nil
		}
	}

	return nil, errors.New("Phone not found.")
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
	fmt.Printf("Starting application....\n")
	router := gin.Default()
	router.GET("/phones", getPhones)
	router.GET("/phones/:id", phoneById)
	router.POST("/phones", createPhone)
	router.PATCH("/checkout", checkoutPhone)
	router.PATCH("/return", returnPhone)
	router.Run("0.0.0.0:8080")
}