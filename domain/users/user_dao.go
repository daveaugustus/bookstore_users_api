package users

import (
	"fmt"

	"github.com/davetweetlive/bookstore_users_api/datasources/mysql/users_db"
	"github.com/davetweetlive/bookstore_users_api/utils/date_utils"
	"github.com/davetweetlive/bookstore_users_api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.CLient.Ping(); err != nil {
		panic(err)
	}
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("The email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("The user %d already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()

	userDB[user.Id] = user
	return nil
}
