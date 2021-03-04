package spotify

import (
	"context"
	"github.com/pkanti/v2/config"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"log"
)

var Auth *clientcredentials.Config
var Client spotify.Client

func init() {
	Auth = &clientcredentials.Config{
		ClientID:     config.Config.Spotify.ID,
		ClientSecret: config.Config.Spotify.Token,
		TokenURL:     spotify.TokenURL,
	}
	token, err := Auth.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	Client = spotify.Authenticator{}.NewClient(token)
	log.Println("Spotify connected")
}
