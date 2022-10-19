package mysql_utils

import (
	"fmt"
	"strings"
	resError "users_api/utils/errors"

	"github.com/go-sql-driver/mysql"
)

const (
	errNoRows = "no rows in result set"
)

func ParseErr(err error) resError.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), errNoRows) {
			return resError.NewNotFoundError("no record matching given id")
		}
		return resError.NewInternalServerError(fmt.Sprintf("error parsing database response %s", err.Error()))
	}

	switch sqlErr.Number {
	case 1062:
		return resError.NewBadRequestError(fmt.Sprintf("duplicate key %s", sqlErr.Message))
	}

	return resError.NewInternalServerError(fmt.Sprintf("error processing request: %s", sqlErr.Message))
}