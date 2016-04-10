package goauth

import(
	"log"
	"github.com/go-ini/ini"
	"golang.org/x/oauth2"
)


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


