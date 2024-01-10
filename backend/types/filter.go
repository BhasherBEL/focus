package types

type Filter struct {
	ID 			int `json:"id"`
	ViewID		int `json:"view_id"`
	TagID		int `json:"tag_id"`
	FilterType	int `json:"filter_type"`
	OptionID	int `json:"option_id"`
}