package config

import (
	"github.com/pelletier/go-toml"
	"github.com/pkanti/v2/database"
	"log"
	"os"
	"path/filepath"
)

type BotConfig struct {
	Discord DiscordConfig     `toml:"Discord"`
	Spotify SpotifyConfig     `toml:"Spotify"`
	Youtube YoutubeConfig     `toml:"Youtube"`
	Db      database.DBConfig `toml:"Database"`
}

type DiscordConfig struct {
	Token string `toml:"API-Token"`
}

type SpotifyConfig struct {
	Token string `toml:"API-Token"`
}

type YoutubeConfig struct {
	Token string `toml:"API-Token"`
}

func LoadAndRefreshConfig(path string) BotConfig {
	outFile, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	loadedConfig := BotConfig{}
	if _, err := os.Stat(outFile); err == nil {
		f, err := os.Open(outFile)
		if err != nil {
			log.Fatal(err)
		}

		err = toml.NewDecoder(f).Decode(&loadedConfig)
		if err != nil {
			log.Fatal(err)
		}

		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}

	outConfig := toml.NewEncoder(f).PromoteAnonymous(true)
	err = outConfig.Encode(loadedConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	return loadedConfig
}
