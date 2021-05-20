package mysql_utils

import (
	"strings"

	"github.com/davetweetlive/bookstore_users_api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record mathcing given id")
		}
		return errors.NewInternalServerError("error parsing databse response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid data")
	}
	return errors.NewInternalServerError("Error processing request")
}
