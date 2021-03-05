package main

import (
	"flag"
	_ "github.com/pkanti/v2/config"
	_ "github.com/pkanti/v2/database"
	_ "github.com/pkanti/v2/discord"
	_ "github.com/pkanti/v2/spotify"
	tests "github.com/pkanti/v2/tests"
)

var runTests = flag.Bool("run-tests", false, "Run tests")

func main() {
	flag.Parse()

	if *runTests {
		tests.TestFindSongYoutube("The Killers", "When You Were Young")
		tests.TestGetMetadataYoutube()
		tests.TestSpotifyGeneral()
		return
	}

}
