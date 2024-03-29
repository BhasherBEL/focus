package types

type View struct {
	ID             int    `json:"id"`
	ProjectID      int    `json:"project_id"`
	Title          string `json:"title"`
	PrimaryTagID   *int   `json:"primary_tag_id"`
	SecondaryTagID *int   `json:"secondary_tag_id"`
	SortTagID      *int   `json:"sort_tag_id"`
	SortDirection  *int   `json:"sort_direction"`
}

type FullView struct {
	ID             int    	`json:"id"`
	ProjectID      int    	`json:"project_id"`
	Title          string 	`json:"title"`
	PrimaryTagID   *int   	`json:"primary_tag_id"`
	SecondaryTagID *int   	`json:"secondary_tag_id"`
	SortTagID      *int   	`json:"sort_tag_id"`
	SortDirection  *int   	`json:"sort_direction"`
	Filters        []Filter `json:"filters"`
}