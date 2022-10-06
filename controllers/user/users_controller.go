package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"status": "Not Ready",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"status": "Not Ready",
	})
}

func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"status": "Not Ready",
	})
}