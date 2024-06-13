package utils

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{})
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
