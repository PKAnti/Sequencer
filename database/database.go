package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	Hostname string `toml:"Hostname"`
	Port     int    `toml:"Port"`
	Database string `toml:"Database"`
	Username string `toml:"Username"`
	Password string `toml:"Password"`
}

func (data *DatabaseConfig) GetURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		data.Username, data.Password, data.Hostname, data.Port, data.Database)
}

func ConnectDatabase(config DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.GetURI())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
