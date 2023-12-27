package db

import (
	"database/sql"
)

var db *sql.DB

func InitDB(driver string, connStr string) error {
	var err error
	db, err = sql.Open(driver, connStr)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS projects (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT
        );
		CREATE TABLE IF NOT EXISTS lists (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			project_id INTEGER,
            title TEXT,
			color TEXT,
			FOREIGN KEY(project_id) REFERENCES projects(id)
        );
		CREATE TABLE IF NOT EXISTS cards (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			list_id INTEGER,
            title TEXT,
			content TEXT,
			FOREIGN KEY(list_id) REFERENCES lists(id)
        );
    `)
	if err != nil {
		return err
	}
	return nil
}
