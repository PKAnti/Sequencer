package youtube

import (
	"context"
	"github.com/pkanti/v2/config"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
)

var Service *youtube.Service

func init() {
	var err error
	ctx := context.Background()
	Service, err = youtube.NewService(ctx,
		option.WithScopes(
			youtube.YoutubeReadonlyScope),
		option.WithAPIKey(config.Config.Youtube.APIKey))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Youtube connected")
}
