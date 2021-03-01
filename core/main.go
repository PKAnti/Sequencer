package main

import (
	//"database/sql"
	"github.com/pkanti/v2/config"
	"github.com/pkanti/v2/database"
	"log"
	"os"
	"path/filepath"
)

func main() {
	regen := false
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i] == "--generate" {
				regen = true
			}
		}
	}

	if regen {
		config.GenerateConfig("config.toml")
		abs, _ := filepath.Abs("config.toml")
		log.Println("config.toml generated at " + abs)
	} else {
		mainConfig := config.LoadConfig("config.toml")
		log.Println("Connecting to " + mainConfig.Db.GetURI())
		db, err := database.ConnectDatabase(mainConfig.Db)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(db)
	}

}
