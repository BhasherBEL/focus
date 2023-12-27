package main

import (
	"database/sql"
)

var db *sql.DB

type Project struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

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
    `)
	if err != nil {
		return err
	}
	return nil
}

func Create(p Project) (int, error) {
	res, err := db.Exec("INSERT INTO projects (title) VALUES (?)", p.Title)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetAll() ([]Project, error) {
	rows, err := db.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var p Project
		if err := rows.Scan(&p.ID, &p.Title); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func Get(id int) (*Project, error) {
	var p Project

	err := db.QueryRow("SELECT * FROM projects WHERE id = ?", id).Scan(&p.ID, &p.Title)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func Delete(id int) error {
	_, err := db.Exec("DELETE FROM projects WHERE id = ?", id)
	return err
}

func Update(p Project) error {
	_, err := db.Exec("UPDATE projects SET title = ? WHERE id = ?", p.Title, p.ID)
	return err
}
