package app

import (
	"users_api/log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Startapp() {
	MapUrls()
	log.Info("about to start the application...")
	router.Run(":8080")
}