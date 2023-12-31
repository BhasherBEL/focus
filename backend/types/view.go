package types

type View struct {
	ID             int    `json:"id"`
	ProjectID      int    `json:"project_id"`
	PrimaryTagID   int    `json:"primary_tag_id"`
	SecondaryTagID int    `json:"secondary_tag_id"`
	Title          string `json:"title"`
}
