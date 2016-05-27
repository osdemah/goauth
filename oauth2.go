package goauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type GOAuth struct {
	Providers map[string]OauthConfig
}

func NewGOAuth() *GOAuth {
	return &GOAuth{Providers: make(map[string]OauthConfig)}
}

func (goauth *GOAuth) AuthHandler(c *gin.Context) {
	oauth, ok := goauth.Providers[c.Param("provider")]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid provider",
		})
		return
	}
	conf := OauthFromStruct(oauth)
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (goauth *GOAuth) CallbackHandler(c *gin.Context) {
	oauth, ok := goauth.Providers[c.Param("provider")]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid provider",
		})
		return
	}
	conf := OauthFromStruct(oauth)
	code := c.Query("code")
	tok, _ := conf.Exchange(oauth2.NoContext, code)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", oauth.ApiURL, nil)
	req.Header.Add("Authorization", "Bearer "+tok.AccessToken)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	profile := Profile{}
	err := json.Unmarshal(body, &profile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, profile)
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
