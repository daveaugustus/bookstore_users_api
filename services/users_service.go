package services

import (
	"github.com/davetweetlive/bookstore_users_api/domain/users"
	"github.com/davetweetlive/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return &user, nil
}
