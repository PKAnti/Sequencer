package tests

import (
	"fmt"
	djspotify "github.com/pkanti/v2/spotify"
	djyoutube "github.com/pkanti/v2/youtube"
	"log"
)

func YoutubeToSpotify(URL string) {
	fmt.Println("Obtaining metadata from " + URL)
	_, metadata, err := djyoutube.GetInfo(URL)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Print("Found on Youtube: ")
	fmt.Println(metadata)

	spotifyURL, spotifyMetadata, err := djspotify.SearchTrackFromMetadata(metadata)
	if err != nil {
		log.Println("Unable to find Spotify mirror for song")
		return
	}
	metadata.Title = spotifyMetadata.Name
	metadata.Artist = make([]string, len(spotifyMetadata.Artists))
	for idx, artist := range spotifyMetadata.Artists {
		metadata.Artist[idx] = artist.Name
	}
	metadata.Album = spotifyMetadata.Album.Name
	fmt.Print("Found on Spotify: ")
	fmt.Println(metadata)
	fmt.Println("URL: " + spotifyURL)
}
