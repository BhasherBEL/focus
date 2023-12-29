package db

import (
	"git.bhasher.com/bhasher/focus/backend/types"
)

func AddTagToCard(ct types.CardTag) error {
	_, err := db.Exec("INSERT INTO cardtags (card_id, tag_id, value) VALUES (?, ?, ?)", ct.CardID, ct.TagID, ct.Value)
	return err
}

func GetAllTagsOfCard(cardID int) ([]types.FullCardTag, error) {
	rows, err := db.Query(`SELECT ct.card_id, ct.tag_id, t.title, ct.value
	FROM cardtags ct
	JOIN tags t ON ct.tag_id = t.id
	WHERE ct.card_id = ?;
	`, cardID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cardtags []types.FullCardTag
	for rows.Next() {
		var ct types.FullCardTag
		if err := rows.Scan(&ct.CardID, &ct.TagID, &ct.TagTitle, &ct.Value); err != nil {
			return nil, err
		}
		cardtags = append(cardtags, ct)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cardtags, nil
}

func DeleteTagOfCard(card_id int, tag_id int) error {
	_, err := db.Exec("DELETE FROM cardtags WHERE card_id = ? AND tag_id = ?", card_id, tag_id)
	return err
}

func DeleteTagsOfCard(card_id int) error {
	_, err := db.Exec("DELETE FROM cardtags WHERE card_id = ?", card_id)
	return err
}

func UpdateTagOfCard(ct types.CardTag) error {
	_, err := db.Exec("UPDATE cardtags SET value = ? WHERE card_id = ? AND tag_id = ?", ct.Value, ct.CardID, ct.TagID)
	return err
}
