package db

import (
	"git.bhasher.com/bhasher/focus/backend/types"
)

func CreateProject(p types.Project) (int, error) {
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

func GetAllProjects() ([]types.Project, error) {
	rows, err := db.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []types.Project
	for rows.Next() {
		var p types.Project
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

func GetProject(id int) (*types.Project, error) {
	var p types.Project

	rows, err := db.Query("SELECT * FROM projects WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	if err := rows.Scan(&p.ID, &p.Title); err != nil {
		return nil, err
	}
	return &p, nil

}

func DeleteProject(id int) (int64, error) {
	res, err := db.Exec("DELETE FROM projects WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateProject(p types.Project) (int64, error) {
	res, err := db.Exec("UPDATE projects SET title = ? WHERE id = ?", p.Title, p.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func ExistProject(id int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM projects WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
