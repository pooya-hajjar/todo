package config

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

type Oauth2 struct {
	Client oauth2.Config
	Ok     bool
}

var AppOath2Config Oauth2

func GoogleConfig() {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	appPort := os.Getenv("APP_PORT")
	redirectURL := fmt.Sprintf("http://localhost:%s/auth/google-callback", appPort)
	scopes := []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"}

	if clientSecret != "" && clientId != "" {
		AppOath2Config.Client = oauth2.Config{
			RedirectURL:  redirectURL,
			ClientID:     clientId,
			ClientSecret: clientSecret,
			Scopes:       scopes,
			Endpoint:     google.Endpoint,
		}
		AppOath2Config.Ok = true
	}
}
