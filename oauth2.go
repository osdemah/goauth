package goauth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"golang.org/x/oauth2"
)

func AuthHandler(c *gin.Context) {
	providers, _ := GetProviders("oauth.ini")
	log.Println(providers)
	log.Println(c.Param("provider"))
	if !StringInSlice(c.Param("provider"), providers) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid provider",
		})
		return
	}
	conf, _ := OauthFromConfig("oauth.ini", c.Param("provider"))
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackHandler(c *gin.Context) {
	providers, _ := GetProviders("oauth.ini")
	if !StringInSlice(c.Param("provider"), providers) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid provider",
		})
		return
	}
	conf, _ := OauthFromConfig("oauth.ini", c.Param("provider"))
	code := c.Query("code")
	tok, _ := conf.Exchange(oauth2.NoContext, code)
	c.JSON(http.StatusOK, gin.H{
		"token": tok.AccessToken,
	})
}

func GetProviders(filepath string) ([]string, error) {
	cfg, err := ini.Load(filepath)
	if err != nil {
		return nil, err
	}
	lists := new(Lists)
	err = cfg.Section("lists").MapTo(lists)
	if err != nil {
		return nil, err
	}
	return lists.Providers, nil
}

func OauthFromStruct(config OauthConfig) oauth2.Config {
	return oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.CallbackURL,
		Scopes:       config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.AuthURL,
			TokenURL: config.TokenURL,
		},
	}
}

func OauthFromConfig(filepath string, provider string) (oauth2.Config, error) {
	cfg, err := ini.Load(filepath)
	if err != nil {
		log.Println(err)
		return oauth2.Config{}, err
	}
	oauth := new(OauthConfig)
	err = cfg.Section(provider).MapTo(oauth)
	if err != nil {
		log.Println(err)
		return oauth2.Config{}, err
	}
	return OauthFromStruct(*oauth), nil
}
