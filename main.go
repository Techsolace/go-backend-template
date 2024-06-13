package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	routes "techsolace/router"
	utils "techsolace/utils"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)

	utils.SetTrustedProxies(router)
	utils.SetupCORS(router)

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	router.Run(":8080")
}
