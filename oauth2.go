package goauth

import(
	"log"
	"github.com/go-ini/ini"
	"golang.org/x/oauth2"
	"net/http"
	"github.com/go-martini/martini"
)

func AuthHandler(res http.ResponseWriter, req *http.Request, params martini.Params){
	providers, err := GetProviders("oauth.ini")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	if !StringInSlice(params["provider"], providers){
		http.Error(res, "invalid provider", 500)
		return
	}
	conf, err := OauthFromConfig("oauth.ini", params["provider"])
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

func CallbackHandler (res http.ResponseWriter, req *http.Request, params martini.Params){
	providers, err := GetProviders("oauth.ini")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	if !StringInSlice(params["provider"], providers){
		http.Error(res, "invalid provider", 500)
		return
	}
	conf, err := OauthFromConfig("oauth.ini", params["provider"])
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	code := req.URL.Query().Get("code")
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	res.Write([]byte(tok.AccessToken))
}

func GetProviders(filepath string) ([]string, error){
	cfg, err := ini.Load(filepath)

	if err != nil{
		log.Println(err)
		return nil, err
	}

	lists := new(Lists)

	err = cfg.Section("lists").MapTo(lists)

	if err != nil{
		log.Println(err)
		return nil, err
	}

	return lists.Providers, nil
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

func OauthFromConfig(filepath string, provider string) (oauth2.Config, error){
	cfg, err := ini.Load(filepath)

	if err != nil{
		log.Println(err)
		return oauth2.Config{}, err
	}

	oauth := new(OauthConfig)

	err = cfg.Section(provider).MapTo(oauth)

	if err != nil{
		log.Println(err)
		return oauth2.Config{}, err
	}

	return OauthFromStruct(*oauth), nil
}


