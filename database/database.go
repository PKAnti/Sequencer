package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase(uri string) (*sql.DB, error) {
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
