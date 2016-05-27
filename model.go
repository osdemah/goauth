package goauth

type Lists struct {
	Providers []string
}

type OauthConfig struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
	AuthURL      string
	TokenURL     string
	ApiURL       string
	Scopes       []string
}

type Profile struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
