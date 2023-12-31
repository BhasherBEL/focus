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
			project_id INTEGER,
            title TEXT,
			content TEXT,
			FOREIGN KEY(project_id) REFERENCES projects(id)
        );
		CREATE TABLE IF NOT EXISTS tags (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			project_id INTEGER,
            title TEXT,
			type int,
			FOREIGN KEY(project_id) REFERENCES projects(id)
        );
		CREATE TABLE IF NOT EXISTS cardtags (
			card_id INTEGER,
			tag_id INTEGER,
			value TEXT,
			PRIMARY KEY(card_id, tag_id),
			FOREIGN KEY(card_id) REFERENCES cards(id)
			FOREIGN KEY(tag_id) REFERENCES tags(id)
        );
		CREATE TABLE IF NOT EXISTS tagsoptions (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			tag_id INTEGER,
			value TEXT,
			FOREIGN KEY(tag_id) REFERENCES tags(id)
        );
    `)
	if err != nil {
		return err
	}
	return nil
}
