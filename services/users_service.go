package services

import (
	"users_api/domain/users"
	crypto_utils "users_api/utils/crypto_utils"
	"users_api/utils/date_utils"
	resError "users_api/utils/errors"
)

type userService struct{}

var (
	UserService userServiceInterface = &userService{}
)

type userServiceInterface interface {
	CreateUser( users.User) (*users.User, *resError.RestError)
	GetUser( int64) (*users.User, *resError.RestError)
	UpdateUser( bool, users.User) (*users.User, *resError.RestError)
	DeleteUser( int64) *resError.RestError
	Search( string) (users.Users,*resError.RestError)
	Login(users.LoginRequest) (*users.User, *resError.RestError)
}

func(s *userService) CreateUser(user users.User) (*users.User, *resError.RestError) {
	if err := user.Validate(); err != nil {
		return nil,err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user,nil
}

func(s *userService) GetUser(userId int64) (*users.User, *resError.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func(s *userService) UpdateUser(isPartial bool,user users.User) (*users.User, *resError.RestError) {
	current, err := s.GetUser(user.Id)
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

func(s *userService) DeleteUser(userId int64) *resError.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}

func(s *userService) Search(status string) (users.Users,*resError.RestError) {
	dao := &users.User{}
	users,err := dao.FindByStatus(status)
	if err != nil {
		return nil,err
	}

	return users,nil
}

func(s *userService) Login(request users.LoginRequest) (*users.User, *resError.RestError) {
	dao := &users.User{
		Email: request.Email,
		Password: crypto_utils.GetMd5(request.Password),
		Status: users.StatusActive,
	}
	if err := dao.FindUserByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao,nil
}