package users

import (
	"users_api/datasource/myql/user_db"
	date_utils "users_api/utils/date_utils"
	resError "users_api/utils/errors"
	mysqlErr "users_api/utils/mysql_utils"

	_ "github.com/go-sql-driver/mysql"
)

const (
	indexUniqueEmail = "user.email_UNIQUE"
	errNoRows = "no rows in result set"
	queryInsertUser ="INSERT INTO user(first_name, last_name, email, date_created) VALUE(?,?,?,?);"
	queryGetUser = "select id, first_name, last_name, email, date_created from user where id=?"
	queryUpdateUser = "update user set first_name=?, last_name=?, email=? where id=?"
	queryDeleteUser = "delete from user where id=?"
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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqlErr.ParseErr(saveErr)
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysqlErr.ParseErr(getErr)
	}

	return nil
}

func(user *User) Update() *resError.RestError {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return resError.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysqlErr.ParseErr(err)
	}

	return nil
}

func (user *User) Delete() *resError.RestError {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return resError.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysqlErr.ParseErr(err)
	}

	return nil

}