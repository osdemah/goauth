package social

import "time"

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

type Github struct {
	Login string `json:"login"`
	ID int `json:"id"`
	AvatarURL string `json:"avatar_url"`
	GravatarID string `json:"gravatar_id"`
	URL string `json:"url"`
	HTMLURL string `json:"html_url"`
	FollowersURL string `json:"followers_url"`
	FollowingURL string `json:"following_url"`
	GistsURL string `json:"gists_url"`
	StarredURL string `json:"starred_url"`
	SubscriptionsURL string `json:"subscriptions_url"`
	OrganizationsURL string `json:"organizations_url"`
	ReposURL string `json:"repos_url"`
	EventsURL string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type string `json:"type"`
	SiteAdmin bool `json:"site_admin"`
	Name string `json:"name"`
	Company string `json:"company"`
	Blog string `json:"blog"`
	Location string `json:"location"`
	Email string `json:"email"`
	Hireable bool `json:"hireable"`
	Bio string `json:"bio"`
	PublicRepos int `json:"public_repos"`
	PublicGists int `json:"public_gists"`
	Followers int `json:"followers"`
	Following int `json:"following"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}