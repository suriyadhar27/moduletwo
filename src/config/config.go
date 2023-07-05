package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "C:/Users/HP/OneDrive/Desktop/rrr/task_one/student.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
