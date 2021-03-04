package tests

import (
	"fmt"
	djspotify "github.com/pkanti/v2/spotify"
	"github.com/zmb3/spotify"
	"log"
	"strings"
)

func TestSpotifyGeneral() {
	artist := "The Killers"
	title := "When You Were Young"
	fmt.Println("Searching Spotify with title=\"" + title + "\" artist=\"" + artist + "\"")

	result, err := djspotify.Client.Search(artist+" "+title, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	var match *spotify.FullTrack
	match = nil
	for _, page := range result.Tracks.Tracks {
		searchArtistMatch := false
		for _, searchArtist := range page.Artists {
			if strings.ToLower(searchArtist.Name) == strings.ToLower(artist) {
				searchArtistMatch = true
				break
			}
		}
		if searchArtistMatch {
			if strings.ToLower(page.Name) == strings.ToLower(title) {
				match = &page
				break
			}
		}
	}

	if match == nil {
		fmt.Println("No qualified match!")
	} else {
		fmt.Print("Found: " + match.Name + " by ")
		for idx, artist := range match.Artists {
			fmt.Print(artist.Name)
			if idx < len(match.Artists)-1 {
				fmt.Print(", ")
			}
		}

		fmt.Printf(" https://open.spotify.com/album/%v?highlight=%v\n", match.Album.ID.String(), match.URI)
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
