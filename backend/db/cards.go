package db

import "git.bhasher.com/bhasher/focus/backend/types"

func CreateCard(c types.Card) (int, error) {
	res, err := db.Exec("INSERT INTO cards (project_id, title, content) VALUES (?, ?, ?)", c.ProjectID, c.Title, c.Content)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetProjectsCards(projectID int) ([]types.FullCard, error) {
	rows, err := db.Query("SELECT * FROM cards WHERE project_id = ?", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []types.FullCard
	for rows.Next() {
		var c types.FullCard
		if err := rows.Scan(&c.ID, &c.ProjectID, &c.Title, &c.Content); err != nil {
			return nil, err
		}

		tags, err := GetCardTags(c.ID, projectID)
		if err != nil {
			return nil, err
		}
		c.Tags = tags

		cards = append(cards, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func GetCard(id int) (*types.FullCard, error) {

	rows, err := db.Query("SELECT * FROM cards WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var c types.FullCard
	rows.Scan(&c.ID, &c.ProjectID, &c.Title, &c.Content)

	tags, err := GetCardTags(id, c.ProjectID)
	if err != nil {
		return nil, err
	}
	c.Tags = tags

	return &c, nil
}

func DeleteCard(id int) (int64, error) {
	res, err := db.Exec("DELETE FROM cards WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateCard(c types.Card) (int64, error) {
	res, err := db.Exec("UPDATE cards SET project_id = ?, title = ?, content = ? WHERE id = ?", c.ProjectID, c.Title, c.Content, c.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func ExistCard(id int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM cards WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
