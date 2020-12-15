package database

import (
	"database/sql"
	"devbook/src/config"
	_ "github.com/go-sql-driver/mysql" // database connection driver
)

func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.ConnectionString)

	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil
}