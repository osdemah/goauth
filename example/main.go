package main

import (
	"github.com/go-martini/martini"
	"github.com/hamed1soleimani/goauth"
)

func main(){
	m := martini.Classic()
	m.Get("/auth/:provider", goauth.AuthHandler)
	m.Get("/auth/:provider/oauth2callback", goauth.CallbackHandler)
	m.Run()

}


