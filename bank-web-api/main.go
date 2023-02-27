package main

import (
	"bank-web-api/configs"
	"bank-web-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()
	routes.AccRoute(router)

	router.Run("localhost:8080")
}
