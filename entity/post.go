package entity

type Post struct {
	ID			string	`json:"id"`
	Text		string	`json:"text"`
	UpdatedAt	int64	`json:"updated_at"`
	CreatedAt	int64	`json:"created_at"`
}
