package types

type CardTag struct {
	CardID   int     `json:"card_id"`
	TagID    int     `json:"tag_id"`
	OptionID *int    `json:"option_id"`
	Value    *string `json:"value"`
}

type FullCardTag struct {
	CardID   int     `json:"card_id"`
	TagID    int     `json:"tag_id"`
	TagTitle string  `json:"tag_title"`
	TagType  int     `json:"tag_type"`
	OptionID *int    `json:"option_id"`
	Value    *string `json:"value"`
}
