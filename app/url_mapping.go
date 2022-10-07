package app

import (
	"users_api/controllers/ping"
	"users_api/controllers/user"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", user.GetUser)
	router.GET("/users", user.SearchUser)
	router.POST("/users", user.CreateUser)
}