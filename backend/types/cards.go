package types

type Card struct {
	ID      int    `json:"id"`
	ListID  int    `json:"list_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
