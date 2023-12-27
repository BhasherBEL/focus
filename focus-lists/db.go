package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type List struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Title     string `json:"title"`
	Color     string `json:"color"`
}

func InitDB(driver string, connStr string) error {
	var err error
	db, err = sql.Open(driver, connStr)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS lists (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			project_id INTEGER,
            title TEXT,
			color TEXT
        );
    `)
	if err != nil {
		return err
	}
	return nil
}

func Create(l List) (int, error) {
	res, err := db.Exec("INSERT INTO lists (project_id, title, color) VALUES (?, ?, ?)", l.ProjectID, l.Title, l.Color)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetAll(projectID int) ([]List, error) {
	rows, err := db.Query("SELECT * FROM lists WHERE project_id = ?", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []List
	for rows.Next() {
		var l List
		if err := rows.Scan(&l.ID, &l.ProjectID, &l.Title, &l.Color); err != nil {
			return nil, err
		}
		lists = append(lists, l)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lists, nil
}

func Get(id int) (*List, error) {
	var l List

	err := db.QueryRow("SELECT * FROM lists WHERE id = ?", id).Scan(&l.ID, &l.ProjectID, &l.Title, &l.Color)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func Delete(id int) error {
	_, err := db.Exec("DELETE FROM lists WHERE id = ?", id)
	return err
}

func Update(l List) error {
	_, err := db.Exec("UPDATE lists SET project_id = ?, title = ?, color = ? WHERE id = ?", l.ProjectID, l.Title, l.Color, l.ID)
	return err
}
