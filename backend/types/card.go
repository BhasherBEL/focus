package types

type Card struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type FullCard struct {
	ID        int           `json:"id"`
	ProjectID int           `json:"project_id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Tags      []FullCardTag `json:"tags"`
}
