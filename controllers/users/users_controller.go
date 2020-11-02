package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/davetweetlive/bookstore_users_api/domain/users"
	"github.com/davetweetlive/bookstore_users_api/services"
	"github.com/davetweetlive/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser takes the request in JSON format in the following formate
// { "id": 123, "first_name": "Dave", "last_name": "Aug", "email": "dave@gmail.com" }
// ReadAll read the body which contains JSON and unmarshal to user var
// Calls the CreateUser function passing user struct as an argument
func CreateUser(c *gin.Context) {
	var user users.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}
