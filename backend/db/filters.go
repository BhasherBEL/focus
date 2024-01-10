package db

import "git.bhasher.com/bhasher/focus/backend/types"

func CreateFilter(filter types.Filter) (int, error) {
	res, err := db.Exec("INSERT INTO filters (view_id, tag_id, filter_type, option_id) VALUES (?, ?, ?, ?)", filter.ViewID, filter.TagID, filter.FilterType, filter.OptionID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetViewFilters(viewID int) ([]types.Filter, error) {
	rows, err := db.Query("SELECT * FROM filters WHERE view_id = ?", viewID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var filters []types.Filter
	for rows.Next() {
		var f types.Filter
		if err := rows.Scan(&f.ID, &f.ViewID, &f.TagID, &f.FilterType, &f.OptionID); err != nil {
			return nil, err
		}

		filters = append(filters, f)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return filters, nil
}

func GetFilter(id int) (*types.Filter, error) {
	rows, err := db.Query("SELECT * FROM filters WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var f types.Filter
	rows.Scan(&f.ID, &f.ViewID, &f.TagID, &f.FilterType, &f.OptionID)

	return &f, nil
}

func UpdateFilter(f types.Filter) (int64, error) {
	res, err := db.Exec("UPDATE filters SET view_id = ?, tag_id = ?, filter_type = ?, option_id = ? WHERE id = ?", f.ViewID, f.TagID, f.FilterType, f.OptionID, f.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteFilter(id int) (int64, error) {
	res, err := db.Exec("DELETE FROM filters WHERE id = ?", id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}