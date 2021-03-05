package core

import "fmt"

type TrackMetadata struct {
	Artist []string
	Album  string
	Title  string
}

func (data TrackMetadata) String() string {
	if len(data.Artist) == 0 {
		return fmt.Sprintf("\"%v\" (Unknown metadata)", data.Title)
	} else if data.Title == "" && data.Album != "" {
		return fmt.Sprintf("\"%v\" by %v", data.Album, data.Artist[0])
	} else if data.Album == "" {
		return fmt.Sprintf("\"%v\" by %v", data.Title, data.Artist[0])
	} else {
		return fmt.Sprintf("\"%v\" by %v (from %v)", data.Title, data.Artist[0], data.Album)
	}
}
