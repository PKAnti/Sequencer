package youtube

import (
	"context"
	"github.com/pkanti/v2/config"
	"github.com/pkanti/v2/core"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"regexp"
	"strings"
)

const (
	typeTrack    = iota
	typeArtist   = iota
	typePlaylist = iota
)

var TrackRegex = regexp.MustCompile("(?:(?:v=)|(?:embed/)|youtu\\.be/)([a-zA-Z0-9_-]*)")
var ChannelRegex = regexp.MustCompile("youtube.com/(?:channel/)([a-zA-Z0-9_-]*)")
var PlaylistRegex = regexp.MustCompile("youtube.com/playlist\\?list=([a-zA-Z0-9_-]*)")

var Service *youtube.Service

func init() {
	var err error
	ctx := context.Background()
	Service, err = youtube.NewService(ctx,
		option.WithScopes(
			youtube.YoutubeReadonlyScope),
		option.WithAPIKey(config.Config.Youtube.Token))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Youtube connected")
}

type InvalidIDError struct {
	msg string
}

func (err InvalidIDError) Error() string {
	return err.msg
}

func GetInfo(URL string) (int, *core.TrackMetadata, error) {
	searchType := -1
	vals := TrackRegex.FindStringSubmatch(URL)
	if len(vals) > 0 {
		searchType = typeTrack
	} else {
		vals = ChannelRegex.FindStringSubmatch(URL)
		if len(vals) > 0 {
			searchType = typeArtist
		} else {
			vals = PlaylistRegex.FindStringSubmatch(URL)
			if len(vals) > 0 {
				searchType = typePlaylist
			}
		}
	}

	key := ""
	if len(vals) > 1 {
		key = vals[1]
	}
	if searchType == typeTrack {
		call := Service.Videos.List([]string{"snippet", "contentDetails", "topicDetails"}).
			Id(key)

		response, err := call.Do()
		if err != nil {
			return -1, nil, err
		} else if len(response.Items) == 0 {
			return -1, nil, InvalidIDError{"No Video Found"}
		} else {
			video := response.Items[0]
			title := video.Snippet.Title
			artist := video.Snippet.ChannelTitle
			album := ""
			if strings.HasSuffix(artist, " - Topic") {
				artist = artist[:len(artist)-8]
			}
			// Youtube music description info
			if strings.HasPrefix(video.Snippet.Description, "Provided to YouTube by") {
				lines := strings.Split(video.Snippet.Description, "\n\n")
				if len(lines) >= 3 {
					data := strings.Split(lines[1], " · ")
					if len(data) >= 1 {
						title = data[0]
						artist = data[1]
					}
					// verify album name in tags
					for _, item := range video.Snippet.Tags {
						if lines[2] == item {
							album = lines[2]
							break
						}
					}
				}
			}
			dataOut := core.TrackMetadata{}
			dataOut.Title = title
			dataOut.Artist = []string{artist}
			dataOut.Album = album
			return searchType, &dataOut, nil
		}
	} else if searchType == typeArtist {
		return -1, nil, InvalidIDError{"Linking artists is not yet supported."}
		/*
			call := Service.Channels.List([]string{"snippet", "contentDetails","topicDetails"}).
				Id(key)

			response, err := call.Do()
			if err != nil {
				return nil, err
			} else if len(response.Items) == 0 {
				return nil, InvalidIDError{"No Channel Found"}
			} else {
				channel := response.Items[0]
				officialArtist := false
				for _, item := range channel.TopicDetails.TopicIds {
					if item == "/m/04rlf" {
						officialArtist = true
						break
					}
				}

				if officialArtist {
					channel.
				} else {

				}
			}
		*/
	} else if searchType == typePlaylist {

		call := Service.Playlists.List([]string{"snippet"}).
			Id(key)

		response, err := call.Do()
		if err != nil {
			return -1, nil, err
		} else if len(response.Items) == 0 {
			return -1, nil, InvalidIDError{"No Video Found"}
		} else {
			albumTitle := response.Items[0].Snippet.Title
			if strings.HasPrefix(albumTitle, "Album - ") {
				albumTitle = albumTitle[8:]
			}
			data := core.TrackMetadata{
				Artist: []string{},
				Album:  albumTitle,
				Title:  "",
			}

			// get first video in playlist
			call := Service.PlaylistItems.List([]string{"snippet"}).
				PlaylistId(key)
			response, err := call.Do()
			if (err == nil) && (response != nil) && (len(response.Items) > 0) {
				_, info, err := GetInfo("youtube.com/watch?v=" + response.Items[0].Snippet.ResourceId.VideoId)
				if err == nil {
					data.Artist = info.Artist
				}
			}

			return searchType, &data, nil
		}
	} else {
		return -1, nil, InvalidIDError{"Could not parse URL"}
	}
}
