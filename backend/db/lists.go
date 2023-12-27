package db

import "git.bhasher.com/bhasher/focus/backend/types"

func CreateList(l types.List) (int, error) {
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

func GetAllListsOf(projectID int) ([]types.List, error) {
	rows, err := db.Query("SELECT * FROM lists WHERE project_id = ?", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []types.List
	for rows.Next() {
		var l types.List
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

func GetList(id int) (*types.List, error) {
	var l types.List

	err := db.QueryRow("SELECT * FROM lists WHERE id = ?", id).Scan(&l.ID, &l.ProjectID, &l.Title, &l.Color)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func DeleteList(id int) error {
	_, err := db.Exec("DELETE FROM lists WHERE id = ?", id)
	return err
}

func UpdateList(l types.List) error {
	_, err := db.Exec("UPDATE lists SET project_id = ?, title = ?, color = ? WHERE id = ?", l.ProjectID, l.Title, l.Color, l.ID)
	return err
}
