package db

import "git.bhasher.com/bhasher/focus/backend/types"

func CreateCard(c types.Card) (int, error) {
	res, err := db.Exec("INSERT INTO cards (list_id, title, content) VALUES (?, ?, ?)", c.ListID, c.Title, c.Content)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetAllCardsOf(listID int) ([]types.Card, error) {
	rows, err := db.Query("SELECT * FROM cards WHERE list_id = ?", listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []types.Card
	for rows.Next() {
		var c types.Card
		if err := rows.Scan(&c.ID, &c.ListID, &c.Title, &c.Content); err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func GetCard(id int) (*types.Card, error) {
	var c types.Card

	err := db.QueryRow("SELECT * FROM cards WHERE id = ?", id).Scan(&c.ID, &c.ListID, &c.Title, &c.Content)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func DeleteCard(id int) error {
	_, err := db.Exec("DELETE FROM cards WHERE id = ?", id)
	return err
}

func UpdateCard(c types.Card) error {
	_, err := db.Exec("UPDATE cards SET list_id = ?, title = ?, content = ? WHERE id = ?", c.ListID, c.Title, c.Content, c.ID)
	return err
}
