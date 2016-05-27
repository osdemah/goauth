package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hamed1soleimani/goauth"
)

func main() {
	r := gin.Default()
	r.GET("/auth/:provider", goauth.AuthHandler)
	r.GET("/auth/:provider/oauth2callback", goauth.CallbackHandler)
	r.Run(":3000")
}
