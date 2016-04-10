package goauth

import(
	"log"
	"github.com/go-ini/ini"
	"golang.org/x/oauth2"
	"net/http"
	"github.com/go-martini/martini"
)

func AuthHandler(res http.ResponseWriter, req *http.Request, params martini.Params){
	if !StringInSlice(params["provider"], GetProviders("oauth.ini")){
		res.Write([]byte("invalid provider"))
		return
	}
	conf := OauthFromConfig("oauth.ini", params["provider"])
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

func CallbackHandler (res http.ResponseWriter, req *http.Request, params martini.Params) (int, string){
	if !StringInSlice(params["provider"], GetProviders("oauth.ini")){
		return 404, "provider not supported"
	}
	conf := OauthFromConfig("oauth.ini", params["provider"])
	code := req.URL.Query().Get("code")
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
		return 401, "Unauthorized"
	}
	return 200, tok.AccessToken
}

func GetProviders(filepath string) []string{
	cfg, err := ini.Load(filepath)

	if err != nil{
		log.Fatal(err)
	}

	lists := new(Lists)

	err = cfg.Section("lists").MapTo(lists)

	if err != nil{
		log.Fatal(err)
	}

	return lists.Providers
}

func OauthFromStruct(config OauthConfig) oauth2.Config{
	return oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL: config.CallbackURL,
		Scopes:       config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.AuthURL,
			TokenURL: config.TokenURL,
		},
	}
}

func OauthFromConfig(filepath string, provider string) oauth2.Config{
	cfg, err := ini.Load(filepath)

	if err != nil{
		log.Fatal(err)
	}

	oauth := new(OauthConfig)

	err = cfg.Section(provider).MapTo(oauth)

	if err != nil{
		log.Fatal(err)
	}

	return OauthFromStruct(*oauth)
}


