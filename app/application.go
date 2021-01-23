package app

import (
	"github.com/davetweetlive/bookstore_users_api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapURLs()

	logger.Info("About to start the application...")
	router.Run(":8081")
}
