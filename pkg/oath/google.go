package oath

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


func GoogleOath() *oauth2.Config{
var GoogleOAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:8080/api/v1/auth/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}
  return GoogleOAuthConfig
}
