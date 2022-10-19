package users

import (
	"fmt"
	"strings"
	"users_api/datasource/myql/user_db"
	"users_api/log"
	resError "users_api/utils/errors"

	_ "github.com/go-sql-driver/mysql"
)

const (
	indexUniqueEmail = "user.email_UNIQUE"
	errNoRows = "no rows in result set"
	queryInsertUser ="INSERT INTO user(first_name, last_name, email, date_created,status, password) VALUE(?,?,?,?,?,?);"
	queryGetUser = "select id, first_name, last_name, email, date_created, status from user where id=?"
	queryUpdateUser = "update user set first_name=?, last_name=?, email=? where id=?"
	queryDeleteUser = "delete from user where id=?"
	queryFindUserByStatus = "select id,first_name,last_name,email,date_created, status from user where status = ?;"
	queryFindUserByEmailAndPassword = "select id, first_name, last_name, email, date_created, status from user where email = ? and password = ? and status = ?;"
)

// var (
// 	usersDB = make(map[int64]*User)
// )

func(user *User) Save() resError.RestError {
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		log.Error("error when trying to prepare save user statement ", err)
		return resError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		log.Error("error when trying to save user ", saveErr)
		return resError.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		log.Error("error when trying to get last insert user id ", err)
		return resError.NewInternalServerError("database error")
	}

	user.Id = userId
	return nil
}

func(user *User) Get() resError.RestError {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		log.Error("error when trying to prepare get user statement ", err)
		return resError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated,  &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), errNoRows) {
			return resError.NewBadRequestError(fmt.Sprintf("not found user with given id %d",user.Id))
		}
		log.Error("error when trying to prepare get user ", getErr)
		return resError.NewInternalServerError("database error")
	}

	return nil
}

func(user *User) Update() resError.RestError {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		log.Error("error when trying to prepare update user statement", err)
		return resError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		log.Error("error when trying to execute update user ", err)
		return resError.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Delete() resError.RestError {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		log.Error("error when trying to prepare delete user statement", err)
		return resError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		log.Error("error when trying to execute delete user ", err)
		return resError.NewInternalServerError("database error")
	}

	return nil

}

func (user *User) FindByStatus(status string) ([]User,resError.RestError){
	stmt, err := user_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		log.Error("error when trying to prepare find user by status statement", err)
		return nil,resError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		log.Error("error when trying to get find user by status ", err)
		return nil, resError.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			log.Error("error when trying to scan user to user struct find user by status user ", err)
			return nil, resError.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, resError.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results,nil
}

func(user *User) FindUserByEmailAndPassword() resError.RestError {
	stmt, err := user_db.Client.Prepare(queryFindUserByEmailAndPassword)
	if err != nil {
		log.Error("error when trying to prepare get user y email and statement ", err)
		return resError.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, user.Status)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated,  &user.Status); getErr != nil {
		if (strings.Contains(getErr.Error(), errNoRows)){
			return resError.NewNotFoundError("invalid login credential")
		}
		log.Error("error when trying to execute get user by email and password ", getErr)
		return resError.NewInternalServerError("database error")
	}

	return nil
}