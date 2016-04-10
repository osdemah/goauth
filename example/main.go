package main

import (
	"log"
	"golang.org/x/oauth2"
	"github.com/go-martini/martini"
	"github.com/hamed1soleimani/goauth"
	"net/http"
)

func main(){
	m := martini.Classic()

	providers := goauth.GetProviders("oauth.ini")

	m.Get("/auth/:provider", func(res http.ResponseWriter, req *http.Request, params martini.Params){
		if !goauth.StringInSlice(params["provider"], providers){
			res.Write([]byte("invalid provider"))
			return
		}
		conf := goauth.OauthFromConfig("oauth.ini", params["provider"])
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		http.Redirect(res, req, url, http.StatusTemporaryRedirect)
	})

	m.Get("/auth/:provider/oauth2callback", func(res http.ResponseWriter, req *http.Request, params martini.Params) (int, string){
		if !goauth.StringInSlice(params["provider"], providers){
			return 404, "provider not supported"
		}
		conf := goauth.OauthFromConfig("oauth.ini", params["provider"])
		code := req.URL.Query().Get("code")
		tok, err := conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Fatal(err)
			return 401, "Unauthorized"
		}
		return 200, tok.AccessToken
	})

	m.Run()

}


