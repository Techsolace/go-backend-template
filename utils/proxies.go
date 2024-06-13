package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetTrustedProxies(router *gin.Engine) {
	proxies := []string{"103.215.227.93"}
	if err := router.SetTrustedProxies(proxies); err != nil {
		log.Fatalf("Error Setting Trusted proxies %v", err)
	}
}
