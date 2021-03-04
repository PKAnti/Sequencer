package config

import (
	"flag"
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
	Token string `toml:"Client-Secret"`
}

type SpotifyConfig struct {
	Token string `toml:"API-Token"`
}

type YoutubeConfig struct {
	Token string `toml:"API-Key"`
}

type DatabaseConfig struct {
	Hostname string `toml:"Hostname"`
	Port     int    `toml:"Port"`
	Database string `toml:"Database"`
	Username string `toml:"Username"`
	Password string `toml:"Password"`
}

var configPath = flag.String("config", "config.toml", ".toml config file path. If missing, the config will be generated.")
var Config = loadConfig()

func init() {
	flag.Parse()
}

func loadConfig() BotConfig {

	// Get absolute path to config
	outFile, err := filepath.Abs(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Generate default config and fill from file
	mainConfig := BotConfig{}
	if _, err := os.Stat(outFile); err == nil {
		f, err := os.Open(outFile)
		if err != nil {
			log.Fatal(err)
		}

		err = toml.NewDecoder(f).Decode(&mainConfig)
		if err != nil {
			log.Fatal(err)
		}

		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Recreate output file
	f, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}

	// Write config back to file
	outConfig := toml.NewEncoder(f).PromoteAnonymous(true)
	err = outConfig.Encode(mainConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Close file
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Configuration loaded")
	return mainConfig
}
