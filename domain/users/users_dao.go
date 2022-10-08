package users

import (
	"fmt"
	"strings"
	"users_api/datasource/myql/user_db"
	date_utils "users_api/utils/date_utils"
	resError "users_api/utils/errors"
)

const (
	indexUniqueEmail = "user.email_UNIQUE"
	errNoRows = "no rows in result set"
	queryInsertUser ="INSERT INTO user(first_name, last_name, email, date_created) VALUE(?,?,?,?);"
	queryGetUser = "select id, first_name, last_name, email, date_created from user where id=?"
)

var (
	usersDB = make(map[int64]*User)
)

func(user *User) Save() *resError.RestError {
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return resError.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return resError.NewBadRequestError("email already exist")
		}
		return resError.NewInternalServerError(err.Error())
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return resError.NewInternalServerError(err.Error())
	}

	user.Id = userId
	return nil
}

func(user *User) Get() *resError.RestError {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		return resError.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errNoRows) {
			return resError.NewNotFoundError(fmt.Sprintf("user %d not found",user.Id))
		}
		fmt.Println(err)
		return resError.NewInternalServerError(fmt.Sprintf("error when try to get user %d: %s ", user.Id, err.Error()))
	}

	return nil
}