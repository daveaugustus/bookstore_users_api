package app

import (
	"github.com/davetweetlive/bookstore_users_api/controllers/ping"
	"github.com/davetweetlive/bookstore_users_api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
}
