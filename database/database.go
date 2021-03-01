package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Hostname string `toml:"Hostname"`
	Port     int    `toml:"Port"`
	Database string `toml:"Database"`
	Username string `toml:"Username"`
	Password string `toml:"Password"`
}

func (data *DBConfig) GetURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		data.Username, data.Password, data.Hostname, data.Port, data.Database)
}

func (data *DBConfig) ValidateNonnull() bool {
	return data.Username != "" &&
		data.Hostname != "" &&
		data.Port > 0 &&
		data.Database != ""
}

func ConnectDatabase(config DBConfig) (*sql.DB, error) {
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
