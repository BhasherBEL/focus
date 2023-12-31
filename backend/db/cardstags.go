package db

import (
	"git.bhasher.com/bhasher/focus/backend/types"
)

func CreateCardTag(ct types.CardTag) error {
	_, err := db.Exec("INSERT INTO cardtags (card_id, tag_id, option_id, value) VALUES (?, ?, ?, ?)", ct.CardID, ct.TagID, ct.OptionID, ct.Value)
	return err
}

func GetCardTags(cardID int, projectID int) ([]types.FullCardTag, error) {
	if projectID < 0 {
		card, err := GetCard(cardID)
		if err != nil {
			return nil, err
		}
		projectID = card.ProjectID
	}

	// rows, err := db.Query(`SELECT t.id, t.title, t.type, COALESCE(ct.option_id, -1), COALESCE(ct.value, '')
	// FROM tags t
	// LEFT JOIN cardtags ct ON ct.tag_id = t.id AND ct.card_id = ?
	// WHERE t.project_id = ?
	// `, cardID, projectID)
	rows, err := db.Query(`SELECT t.id, t.title, t.type, ct.option_id, ct.value
	FROM cardtags ct
	INNER JOIN tags t ON t.id = ct.tag_id
	WHERE ct.card_id = ? AND t.project_id = ?
	`, cardID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cardtags []types.FullCardTag
	for rows.Next() {
		ct := types.FullCardTag{CardID: cardID}
		if err := rows.Scan(&ct.TagID, &ct.TagTitle, &ct.TagType, &ct.OptionID, &ct.Value); err != nil {
			return nil, err
		}
		cardtags = append(cardtags, ct)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cardtags, nil
}

func DeleteCardTag(card_id int, tag_id int) (int64, error) {
	res, err := db.Exec("DELETE FROM cardtags WHERE card_id = ? AND tag_id = ?", card_id, tag_id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DeleteCardTags(card_id int) (int64, error) {
	res, err := db.Exec("DELETE FROM cardtags WHERE card_id = ?", card_id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateCardTag(ct types.CardTag) (int64, error) {
	res, err := db.Exec("UPDATE cardtags SET option_id = ?, value = ? WHERE card_id = ? AND tag_id = ?", ct.OptionID, ct.Value, ct.CardID, ct.TagID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
