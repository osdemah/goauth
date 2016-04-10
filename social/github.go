package social

import(
	"log"
	"io/ioutil"
	"net/http"

	"github.com/go-ini/ini"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/oauth2"
	goauth2 "golang.org/x/oauth2"
	"fmt"
	"encoding/json"
)

func GithubOauthStruct(config OauthConfig) martini.Handler{
	oauth2.PathCallback = config.CallbackURL
	oauth2.PathError = config.ErrorURL
	oauth2.PathLogin = config.LoginURL
	oauth2.PathLogout = config.LogoutURL

	return oauth2.Github(
		&goauth2.Config{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			Scopes:       config.Scopes,
			RedirectURL:  config.RedirectURL,
		},
	)
}

func GithubOauthConfig(filepath string) martini.Handler{
	cfg, err := ini.Load(filepath)

	if err != nil{
		log.Fatal(err)
	}

	oauth := new(OauthConfig)

	err = cfg.Section("github").MapTo(oauth)

	if err != nil{
		log.Fatal(err)
	}

	return GithubOauthStruct(*oauth)
}

func GithubProfileJson(token string) string{
	response, err := http.Get("https://api.github.com/user?access_token=" + token)
	fmt.Println("https://api.github.com/user?access_token=" + token)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}

func GithubProfileStruct(token string) Github{
	github := new(Github)
	json.Unmarshal([]byte(GoogleProfileJson(token)), &*github)
	return *github
}



