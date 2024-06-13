package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetTrustedProxies(router *gin.Engine) {
	proxies := []string{"0.0.0.0"}
	if err := router.SetTrustedProxies(proxies); err != nil {
		log.Fatalf("Error Setting Trusted proxies %v", err)
	}
}
