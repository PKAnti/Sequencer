package config

import (
	"github.com/pelletier/go-toml"
	"log"
	"os"
	"path/filepath"
)

type BotConfig struct {
	Discord DiscordConfig  `toml:"Discord"`
	Spotify SpotifyConfig  `toml:"Spotify"`
	Youtube YoutubeConfig  `toml:"Youtube"`
	Db      DatabaseConfig `toml:"Database"`
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

type DatabaseConfig struct {
	Uri string `toml:"URI"`
}

func GenerateConfig(path string, overwrite bool) {
	outFile, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(outFile); err == nil && overwrite == false {
		log.Fatal("Config file to generate already exists.")
	}

	f, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}

	sampleConfig := BotConfig{}

	outConfig := toml.NewEncoder(f).PromoteAnonymous(true)
	err = outConfig.Encode(sampleConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func LoadConfig(path string) BotConfig {
	config, err := toml.LoadFile(path)
	if err != nil {
		log.Fatal("Error reading config! Generate a basic config with \"" + os.Args[0] + " --generate" + "\nError: " + err.Error())
	}

	read_config := BotConfig{}
	err = config.Unmarshal(&read_config)
	if err != nil {
		log.Fatal(err)
	}
	return read_config
}
