package db

import "git.bhasher.com/bhasher/focus/backend/types"

func CreateView(v types.View) (int, error) {
	res, err := db.Exec("INSERT INTO views (project_id, primary_tag_id, secondary_tag_id, title, sort_tag_id, sort_direction) VALUES (?, ?, ?, ?, ?, ?)", v.ProjectID, v.PrimaryTagID, v.SecondaryTagID, v.Title, v.SortTagID, v.SortDirection)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetProjectViews(projectID int) ([]types.FullView, error) {
	rows, err := db.Query("SELECT * FROM views WHERE project_id = ?", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []types.FullView
	for rows.Next() {
		var v types.FullView
		if err := rows.Scan(&v.ID, &v.ProjectID, &v.PrimaryTagID, &v.SecondaryTagID, &v.Title, &v.SortTagID, &v.SortDirection); err != nil {
			return nil, err
		}

		filters, err := GetViewFilters(v.ID)
		if err != nil {
			return nil, err
		}
		v.Filters = filters

		views = append(views, v)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return views, nil
}

func GetView(id int) (*types.FullView, error) {
	rows, err := db.Query("SELECT * FROM views WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var v types.FullView
	rows.Scan(&v.ID, &v.ProjectID, v.PrimaryTagID, v.SecondaryTagID, v.Title, v.SortTagID, v.SortDirection)

	filters, err := GetViewFilters(id)
	if(err != nil) {
		return nil, err
	}
	v.Filters = filters

	return &v, nil
}

func UpdateView(v types.View) (int64, error) {
	res, err := db.Exec("UPDATE views SET project_id = ?, primary_tag_id = ?, secondary_tag_id = ?, title = ?, sort_tag_id = ?, sort_direction = ? WHERE id = ?", v.ProjectID, v.PrimaryTagID, v.SecondaryTagID, v.Title, v.SortTagID, v.SortDirection, v.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteView(id int) (int64, error) {
	res, err := db.Exec("DELETE FROM views WHERE id = ?", id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func ExistView(id int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM views WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
