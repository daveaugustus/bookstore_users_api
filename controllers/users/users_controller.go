package users

import (
	"net/http"

	"github.com/davetweetlive/bookstore_users_api/domain/users"
	"github.com/davetweetlive/bookstore_users_api/services"
	"github.com/davetweetlive/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

//
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

//
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
