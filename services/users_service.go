package services

import (
	"users_api/domain/users"
	resError "users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *resError.RestError) {

	return &user,nil
}