package main

import (
	"github.com/pkanti/v2/config"
	"github.com/pkanti/v2/database"
	"log"
	"os"
)

func main() {
	regen := false
	overwrite := false
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i] == "--generate" {
				regen = true
			}
			if os.Args[i] == "--overwrite" {
				overwrite = true
			}
		}
	}

	if regen {
		config.GenerateConfig("config.toml", overwrite)
	} else {
		mainConfig := config.LoadConfig("config.toml")
		db, err := database.ConnectDatabase(mainConfig.Db.Uri)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(db)
	}

}
