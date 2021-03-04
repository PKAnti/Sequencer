package youtube

import (
	"fmt"
	"github.com/pkanti/v2/youtube"
	ytapi "google.golang.org/api/youtube/v3"
	"log"
	"strings"
	"time"
)

func TestFindSongYoutube(artist, title string) {

	fmt.Printf("Searching for: \"%v\" by %v\n", title, artist)

	call := youtube.Service.Search.List([]string{"id", "snippet"}).
		Q(artist).
		MaxResults(5).
		Type("channel")

	response, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	channelID := ""
	for _, item := range response.Items {
		fmt.Printf("[%v] %v\n", item.Id.ChannelId, item.Snippet.Title)
		id := item.Id.ChannelId
		title := item.Snippet.Title
		if len(title) > 8 &&
			title[len(title)-8:] == " - Topic" {
			channelID = id
			break
		}
	}

	if channelID == "" {
		fmt.Println("No matching artist found")
		return
	} else {
		fmt.Println("Found channel ID: " + channelID)
	}

	call = youtube.Service.Search.List([]string{"id", "snippet"}).
		Q(title).
		MaxResults(5).
		ChannelId(channelID)

	response, err = call.Do()
	if err != nil {
		log.Fatal(err)
	}

	var VideoSnippet *ytapi.SearchResultSnippet
	VideoSnippet = nil
	VideoID := ""
	EarliestPublish := time.Unix(1<<63-62135596801, 999999999)
	for _, item := range response.Items {
		fmt.Printf("[%v] %v\n", item.Id.VideoId, item.Snippet.Title)
		id := item.Id.VideoId
		snippet := item.Snippet
		strTime, _ := time.Parse(time.RFC3339, snippet.PublishedAt)
		if strings.ToLower(snippet.Title) == strings.ToLower(title) &&
			strTime.Before(EarliestPublish) {
			VideoID = id
			VideoSnippet = snippet
			EarliestPublish = strTime
		}
	}

	if VideoSnippet == nil {
		fmt.Println("No suitable video found")
	} else {
		fmt.Println("Found: http://www.youtube.com/watch?v=" + VideoID)
		fmt.Println("Channel: " + VideoSnippet.ChannelTitle)
		fmt.Println("Title: " + VideoSnippet.Title)
	}
}
