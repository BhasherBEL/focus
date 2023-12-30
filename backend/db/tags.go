package db

import "git.bhasher.com/bhasher/focus/backend/types"

func CreateTag(t types.Tag) (int, error) {
	res, err := db.Exec("INSERT INTO tags (project_id, title, type) VALUES (?, ?, ?)", t.ProjectID, t.Title, t.Type)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetProjectTags(projectID int) ([]types.Tag, error) {
	rows, err := db.Query("SELECT * FROM tags WHERE project_id = ?", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []types.Tag
	for rows.Next() {
		var t types.Tag
		if err := rows.Scan(&t.ID, &t.ProjectID, &t.Title, &t.Type); err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func GetTag(id int) (*types.Tag, error) {
	var t types.Tag

	rows, err := db.Query("SELECT * FROM tags WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	rows.Scan(&t.ID, &t.ProjectID, &t.Title, &t.Type)

	return &t, nil
}

func DeleteTag(id int) (int64, error) {
	res, err := db.Exec("DELETE FROM tags WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateTag(t types.Tag) (int64, error) {
	res, err := db.Exec("UPDATE tags SET project_id = ?, title = ?, type = ? WHERE id = ?", t.ProjectID, t.Title, t.Type, t.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func ExistTag(id int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tas WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
