package main

import (
	"flag"
	"fmt"
	_ "github.com/pkanti/v2/config"
	ytapi "google.golang.org/api/youtube/v3"
	"strings"
	"time"

	_ "github.com/pkanti/v2/database"
	_ "github.com/pkanti/v2/discord"
	"github.com/pkanti/v2/youtube"
	"log"
)

var (
	artist = flag.String("artist", "The Killers", "Artist")
	title  = flag.String("title", "When You Were Young", "Song title")
)

func main() {
	flag.Parse()

	fmt.Printf("Searching for: \"%v\" by %v\n", *title, *artist)

	call := youtube.Service.Search.List([]string{"id", "snippet"}).
		Q(*artist).
		MaxResults(5).
		Type("channel")

	response, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	channels := make(map[string]*ytapi.SearchResultSnippet)

	for _, item := range response.Items {
		channels[item.Id.ChannelId] = item.Snippet
	}

	printTitleMap(channels)

	channelID := ""
	for id, snippet := range channels {
		if len(snippet.ChannelTitle) > 8 &&
			snippet.ChannelTitle[len(snippet.ChannelTitle)-8:] == " - Topic" {
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
		Q(*title).
		MaxResults(5).
		ChannelId(channelID)

	response, err = call.Do()
	if err != nil {
		log.Fatal(err)
	}

	videos := make(map[string]*ytapi.SearchResultSnippet)

	for _, item := range response.Items {
		videos[item.Id.VideoId] = item.Snippet
	}

	printTitleMap(videos)

	var VideoSnippet *ytapi.SearchResultSnippet
	VideoSnippet = nil
	VideoID := ""
	EarliestPublish := time.Unix(1<<63-62135596801, 999999999)
	for id, snippet := range videos {
		strTime, _ := time.Parse(time.RFC3339, snippet.PublishedAt)
		if strings.ToLower(snippet.Title) == strings.ToLower(*title) &&
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

func printTitleMap(data map[string]*ytapi.SearchResultSnippet) {
	for id, snippet := range data {
		fmt.Printf("[%v] %v\n", id, snippet.Title)
	}
}
