package user

import (
	"net/http"
	"strconv"
	"users_api/domain/users"
	"users_api/services"

	resError "users_api/utils/errors"

	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *resError.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)

	if userErr != nil {
		return 0, resError.NewBadRequestError("user id should be a number")
	}

	return userId, nil
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		resErr := resError.NewBadRequestError("user id should be a number")
		c.JSON(resErr.Status, resErr)
		return
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return 
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
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

	result, err := services.UserService.CreateUser(user)
	if err != nil {
		//TODO: handle error
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func SearchUser(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		resErr := resError.NewBadRequestError("user id should be a number")
		c.JSON(resErr.Status, resErr)
		return
	}

	var user users.User

	//This function is same as we Readall and unmarshall the request
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: handle json error
		var resErr resError.RestError
		resErr = *resError.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, resErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method== http.MethodPatch

	result, err := services.UserService.UpdateUser(isPartial,user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status":"deleted"})
}
