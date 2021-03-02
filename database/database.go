package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkanti/v2/config"
	"log"
)

func getURI(data config.DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		data.Username, data.Password, data.Hostname, data.Port, data.Database)
}

func ValidateNonnull(data config.DatabaseConfig) bool {
	return data.Username != "" &&
		data.Hostname != "" &&
		data.Port > 0 &&
		data.Database != ""
}

var Database *sql.DB

func init() {
	if !ValidateNonnull(config.Config.Db) {
		log.Fatal("Invalid database configuration")
	}

	var err error
	Database, err = sql.Open("mysql", getURI(config.Config.Db))
	if err != nil {
		log.Fatal(err)
	}

	err = Database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
}
