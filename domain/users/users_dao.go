package users

import (
	"fmt"
	resError "users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func(user *User) Save() *resError.RestError {
	currentUser := usersDB[user.Id]

	if currentUser != nil {
		if currentUser.Email == user.Email {
			return resError.NewBadRequestError(fmt.Sprintf("email %s id already registered", user.Email))
		}
		
		return resError.NewBadRequestError(fmt.Sprintf("user %d is already exist", user.Id))
	}

	usersDB[user.Id] = user

	return nil
}

func(user *User) Get() *resError.RestError {
	result := usersDB[user.Id]

	if result == nil {
		return resError.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}