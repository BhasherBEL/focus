package types

type Tag struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Title     string `json:"title"`
	Type      int    `json:"type"`
}
