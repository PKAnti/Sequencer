package main

import (
	"flag"
	"github.com/pkanti/v2/config"
	_ "github.com/pkanti/v2/database"
	_ "github.com/pkanti/v2/discord"
	_ "github.com/pkanti/v2/spotify"
	tests "github.com/pkanti/v2/tests"
)

func main() {
	flag.Parse()

	if *config.RunTests {
		tests.TestFindSongYoutube("The Killers", "When You Were Young")
		tests.TestFindSongYoutube("Sir Sly", "&Run")
		tests.TestGetMetadataYoutube()
		tests.TestSpotifyGeneral()
		tests.YoutubeToSpotify("https://music.youtube.com/watch?v=2Vu0WhZOvAQ&feature=share")
		return
	}

}
