package tests

import (
	"fmt"
	"github.com/pkanti/v2/core"
	djspotify "github.com/pkanti/v2/spotify"
	"github.com/zmb3/spotify"
	"log"
	"strings"
)

func TestSpotifyGeneral() {
	artist := "The Killers"
	title := "When You Were Young"
	fmt.Println("Searching Spotify with title=\"" + title + "\" artist=\"" + artist + "\"")

	localData := core.TrackMetadata{
		Artist: []string{artist},
		Title:  title,
		Album:  "",
	}

	spotifyURL, metadata, err := djspotify.SearchTrackFromMetadata(&localData)
	if err == nil {
		fmt.Println("No qualified match!")
	} else {
		fmt.Print("Found: " + metadata.Name + " by ")
		for idx, artist := range metadata.Artists {
			fmt.Print(artist.Name)
			if idx < len(metadata.Artists)-1 {
				fmt.Print(", ")
			}
		}

		fmt.Println(" " + spotifyURL)
	}

	playlistURL := "https://open.spotify.com/playlist/7GRdKkdbUprchkrf0hVqqX?si=cKUs9sn2QJ6X3NUiD1S33g"
	fmt.Println("Testing playlist identifier on " + playlistURL)
	var key string
	_, err = fmt.Sscanf(playlistURL, "https://open.spotify.com/playlist/%v", &key)
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(key, "?") {
		key = key[:strings.Index(key, "?")]
	}

	foundPlaylist, err := djspotify.Client.GetPlaylist(spotify.ID(key))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Playlist: \"%v\" (Tracks: %d)\n", foundPlaylist.Name, len(foundPlaylist.Tracks.Tracks))

	albumURL := "https://open.spotify.com/album/0UoIEoMO8GchrDZQRSIzmI?si=oCBloKJ_QX6FB3uOvVUrcg"
	fmt.Println("Testing album identifier on " + albumURL)
	var albumKey string
	_, err = fmt.Sscanf(albumURL, "https://open.spotify.com/album/%v", &albumKey)
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(albumKey, "?") {
		albumKey = albumKey[:strings.Index(albumKey, "?")]
	}

	foundAlbum, err := djspotify.Client.GetAlbum(spotify.ID(albumKey))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album: \"%v\" by %v (Tracks: %d)", foundAlbum.Name, foundAlbum.Artists[0].Name, len(foundAlbum.Tracks.Tracks))

}
