package social

type OauthConfig struct {
	ClientID string
	ClientSecret string
	RedirectURL string
	LoginURL string
	LogoutURL string
	ErrorURL string
	CallbackURL string
	Scopes []string
}

type Google struct {
	ID string `json:"id"`
	Email string `json:"email"`
	VerifiedEmail bool `json:"verified_email"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Link string `json:"link"`
	Picture string `json:"picture"`
	Gender string `json:"gender"`
	Locale string `json:"locale"`
}