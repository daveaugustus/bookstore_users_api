package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	if err := router.Run(":8900"); err != nil {
		fmt.Println("Couldn't start the server at port 8080")
	}
}
