package services

import (
	"users_api/domain/users"
	resError "users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *resError.RestError) {
	if err := user.Validate(); err != nil {
		return nil,err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user,nil
}

func GetUser(userId int64) (*users.User, *resError.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}