package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	routes "techsolace/router"
	utils "techsolace/utils"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.SetupRoutes(router)

	utils.SetTrustedProxies(router)
	utils.SetupCORS(router)

	router.NoRoute(utils.HandleNotFound)

	port := utils.GetPort()
	fmt.Printf("Listening to %v", port)
	router.Run(":" + port)
}
