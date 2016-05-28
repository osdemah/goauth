package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hamed1soleimani/goauth"
)

func OauthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		//customize your needs in the middleware
		//user profile and token are in the context variables
		c.JSON(http.StatusOK, c.Value("profile"))
	}
}

func main() {
	r := gin.Default()
	auth := goauth.NewGOAuth()
	auth.Providers["google"] = goauth.OauthConfig{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_SECRET",
		CallbackURL:  "http://127.0.0.1:3000/auth/google/oauth2callback",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://accounts.google.com/o/oauth2/token",
		ApiURL:       "https://www.googleapis.com/oauth2/v2/userinfo?fields=email%2Cname%2Cpicture",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
	}

	r.GET("/auth/:provider", auth.AuthHandler)
	r.GET("/auth/:provider/oauth2callback", OauthMiddleware(), auth.CallbackHandler)
	r.Run(":3000")
}
