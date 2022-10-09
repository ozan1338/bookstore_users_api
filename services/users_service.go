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

func UpdateUser(isPartial bool,user users.User) (*users.User, *resError.RestError) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil,err
	}

	if  !isPartial {
		current.FirstName = user.FirstName
		current.Email = user.Email
		current.LastName = user.LastName
	} else {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	}



	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userId int64) *resError.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}