package types

type Tag struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Title     string `json:"title"`
	Type      int    `json:"type"`
}

type TagOption struct {
	ID    int    `json:"id"`
	TagID int    `json:"tag_id"`
	Value string `json:"value"`
}

type FullTag struct {
	ID        int         `json:"id"`
	ProjectID int         `json:"project_id"`
	Title     string      `json:"title"`
	Type      int         `json:"type"`
	Options   []TagOption `json:"options"`
}
