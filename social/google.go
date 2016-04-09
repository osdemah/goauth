package social

import(
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-ini/ini"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/oauth2"
	goauth2 "golang.org/x/oauth2"
)

func GoogleOauthStruct(config OauthConfig) martini.Handler{
	return oauth2.Google(
		&goauth2.Config{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			Scopes:       config.Scopes,
			RedirectURL:  config.RedirectURL,
		},
	)
}

func GoogleOauthConfig(filepath string) martini.Handler{
	cfg, err := ini.Load(filepath)

	if err != nil{
		log.Fatal(err)
	}

	oauth := new(OauthConfig)

	err = cfg.Section("google").MapTo(oauth)

	if err != nil{
		log.Fatal(err)
	}

	return GoogleOauthStruct(*oauth)
}

func GoogleProfileStruct(token string) Google{
	google := new(Google)
	json.Unmarshal([]byte(GoogleProfileJson(token)), &*google)
	return *google
}

func GoogleProfileJson(token string) string{
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
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



