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

func GetProjectTags(projectID int) ([]types.FullTag, error) {
	rows, err := db.Query("SELECT * FROM tags WHERE project_id = ?", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []types.FullTag
	for rows.Next() {
		var t types.FullTag
		if err := rows.Scan(&t.ID, &t.ProjectID, &t.Title, &t.Type); err != nil {
			return nil, err
		}

		t.Options, err = GetTagOptions(t.ID)
		if err != nil {
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
	err := db.QueryRow("SELECT COUNT(*) FROM tags WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func CreateTagOption(to types.TagOption) (int, error) {
	res, err := db.Exec("INSERT INTO tagsoptions (tag_id, value) VALUES (?, ?)", to.TagID, to.Value)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetTagOptions(tagID int) ([]types.TagOption, error) {
	rows, err := db.Query("SELECT * FROM tagsoptions WHERE tag_id = ?", tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var options []types.TagOption
	for rows.Next() {
		var to types.TagOption
		if err := rows.Scan(&to.ID, &to.TagID, &to.Value); err != nil {
			return nil, err
		}
		options = append(options, to)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return options, nil
}

func DeleteTagOption(id int) (int64, error) {
	res, err := db.Exec("DELETE FROM tagsoptions WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateTagOption(to types.TagOption) (int64, error) {
	res, err := db.Exec("UPDATE tagsoptions SET tag_id = ?, value = ? WHERE id = ?", to.TagID, to.Value, to.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DeleteTagOptions(tagID int) error {
	_, err := db.Exec("DELETE FROM tagsoptions WHERE tag_id = ?", tagID)
	return err
}
