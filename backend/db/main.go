package db

import (
	"database/sql"
)

var db *sql.DB

const DB_VERSION = 1

func InitDB(driver string, connStr string) error {
	var err error
	db, err = sql.Open(driver, connStr)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_version (
			version INTEGER PRIMARY KEY,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
        CREATE TABLE IF NOT EXISTS projects (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT
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
			option_id TEXT,
			value TEXT,
			PRIMARY KEY(card_id, tag_id),
			FOREIGN KEY(card_id) REFERENCES cards(id)
			FOREIGN KEY(tag_id) REFERENCES tags(id)
			FOREIGN KEY(option_id) REFERENCES tagsoptions(id)
        );
		CREATE TABLE IF NOT EXISTS tagsoptions (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			tag_id INTEGER,
			value TEXT,
			FOREIGN KEY(tag_id) REFERENCES tags(id)
        );
		CREATE TABLE IF NOT EXISTS views (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			project_id INTEGER,
			primary_tag_id INTEGER,
			secondary_tag_id INTEGER,
			title TEXT,
			sort_tag_id INTEGER,
			sort_direction INTEGER,
			FOREIGN KEY(project_id) REFERENCES projects(id),
			FOREIGN KEY(primary_tag_id) REFERENCES tags(id),
			FOREIGN KEY(secondary_tag_id) REFERENCES tags(id)
		);

		INSERT INTO schema_version (version)
		SELECT ? WHERE NOT EXISTS (SELECT 1 FROM schema_version);
    `, DB_VERSION)

	if err != nil {
		return err
	}

	return updateDB()
}

func updateDB() error {

	var currentVersion int

	err := db.QueryRow("SELECT MAX(version) FROM schema_version").Scan(&currentVersion)
	if err != nil {
		return err
	}

	if currentVersion < 2 {
		_, err := db.Exec(`
			ALTER TABLE views 
			ADD COLUMN sort_tag_id INTEGER;
			ALTER TABLE views 
			ADD COLUMN sort_direction INTEGER;
			
			INSERT INTO schema_version (version)
			SELECT ? WHERE NOT EXISTS (SELECT 1 FROM schema_version WHERE version = ?);
    	`, 2, 2)
		if err != nil {
			return err
		}
	}

	return nil
}
