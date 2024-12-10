package dropbox

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var DropboxConfig *oauth2.Config

func init() {

	err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }
	clientID := os.Getenv("DROPBOX_CLIENT_ID")
	clientSecret := os.Getenv("DROPBOX_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Fatal("DROPBOX_CLIENT_ID or DROPBOX_CLIENT_SECRET is not set in environment variables")
	}

	DropboxConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/api/auth/callback",
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.dropbox.com/oauth2/authorize",
			TokenURL: "https://api.dropbox.com/oauth2/token",
		},
	}
}