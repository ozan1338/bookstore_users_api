package user

import (
	"fmt"
	"net/http"
	"users_api/domain/users"
	"users_api/services"

	resError "users_api/utils/errors"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"status": "Not Ready",
	})
}

func CreateUser(c *gin.Context) {
	var user users.User

	//This function is same as we Readall and unmarshall the request
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: handle json error
		var resErr resError.RestError
		resErr = *resError.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, resErr)
		return
	}

	fmt.Println(user)

	result, err := services.CreateUser(user)
	if err != nil {
		//TODO: handle error
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"status": "Not Ready",
	})
}