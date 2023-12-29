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

func GetAllCardsOf(projectID int) ([]types.FullCard, error) {
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

		tags, err := GetAllTagsOfCard(c.ID)
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
	var c types.FullCard

	err := db.QueryRow("SELECT * FROM cards WHERE id = ?", id).Scan(&c.ID, &c.ProjectID, &c.Title, &c.Content)
	if err != nil {
		return nil, err
	}

	tags, err := GetAllTagsOfCard(id)
	if err != nil {
		return nil, err
	}
	c.Tags = tags

	return &c, nil
}

func DeleteCard(id int) error {
	_, err := db.Exec("DELETE FROM cards WHERE id = ?", id)
	return err
}

func UpdateCard(c types.Card) error {
	_, err := db.Exec("UPDATE cards SET project_id = ?, title = ?, content = ? WHERE id = ?", c.ProjectID, c.Title, c.Content, c.ID)
	return err
}
