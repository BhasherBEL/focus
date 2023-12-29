package db

import (
	"git.bhasher.com/bhasher/focus/backend/types"
)

func AddTagToCard(ct types.CardTag) error {
	_, err := db.Exec("INSERT INTO cardtags (card_id, tag_id, value) VALUES (?, ?, ?)", ct.CardID, ct.TagID, ct.Value)
	return err
}

func GetAllTagsOfCard(cardID int, projectID int) ([]types.FullCardTag, error) {
	if projectID < 0 {
		card, err := GetCard(cardID)
		if err != nil {
			return nil, err
		}
		projectID = card.ProjectID
	}

	rows, err := db.Query(`SELECT t.id, t.title, COALESCE(ct.value, '')
	FROM tags t
	LEFT JOIN cardtags ct ON ct.tag_id = t.id AND ct.card_id = ?
	WHERE t.project_id = ?
	`, cardID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cardtags []types.FullCardTag
	for rows.Next() {
		ct := types.FullCardTag{CardID: cardID}
		if err := rows.Scan(&ct.TagID, &ct.TagTitle, &ct.Value); err != nil {
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
