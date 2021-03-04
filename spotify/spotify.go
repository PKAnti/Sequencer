package spotify

import (
	"context"
	"fmt"
	"github.com/pkanti/v2/config"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"log"
)

var auth *clientcredentials.Config

func init() {
	auth = &clientcredentials.Config{
		ClientID:     config.Config.Spotify.ID,
		ClientSecret: config.Config.Spotify.Token,
		TokenURL:     spotify.TokenURL,
	}
	token, err := auth.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	msg, page, err := client.FeaturedPlaylists()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	for _, playlist := range page.Playlists {
		fmt.Println(" ", playlist.Name)
	}
}
