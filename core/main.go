package main

import (
	//"database/sql"
	"github.com/pkanti/v2/config"
	"github.com/pkanti/v2/database"
	"log"
)

func main() {
	mainConfig := config.LoadAndRefreshConfig("config.toml")
	validConfig := true
	if !mainConfig.Db.ValidateNonnull() {
		log.Println("Invalid database connection details")
		validConfig = false
	}
	if !validConfig {
		log.Fatal("Unable to continue due to previous errors.")
	}

	log.Println("Connecting to " + mainConfig.Db.GetURI())
	db, err := database.ConnectDatabase(mainConfig.Db)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(db)

}
