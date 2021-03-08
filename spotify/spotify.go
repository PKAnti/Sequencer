package spotify

import (
	"context"
	"fmt"
	"github.com/pkanti/v2/config"
	"github.com/pkanti/v2/core"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"strings"
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

type InvalidIDError struct {
	msg string
}

func (err InvalidIDError) Error() string {
	return err.msg
}

func SearchTrackFromMetadata(metadata *core.TrackMetadata) (string, *spotify.FullTrack, error) {
	if metadata.Title == "" {
		return "", nil, InvalidIDError{"Title is required for track searches"}
	}
	if len(metadata.Artist) == 0 {
		return "", nil, InvalidIDError{"Artist is required for track searches"}
	}

	fmt.Println("Searching Spotify with title=\"" + metadata.Title + "\" artist=\"" + metadata.Artist[0] + "\" album=\"" + metadata.Album + "\"")

	result, err := Client.Search(metadata.Title+" "+metadata.Artist[0]+" "+metadata.Album, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	var match *spotify.FullTrack
	match = nil
	for _, page := range result.Tracks.Tracks {
		searchArtistMatch := false
		for _, searchArtist := range page.Artists {
			if strings.ToLower(searchArtist.Name) == strings.ToLower(metadata.Artist[0]) {
				searchArtistMatch = true
				break
			}
		}
		if searchArtistMatch {
			if strings.ToLower(page.Name) == strings.ToLower(metadata.Title) {
				match = &page
				break
			}
		}
	}

	if match == nil {
		return "", nil, InvalidIDError{"No valid match found"}
	} else {
		return match.SimpleTrack.ExternalURLs["spotify"], match, nil
	}
}
