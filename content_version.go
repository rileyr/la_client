package littleaspen

import "time"

type ContentVersion struct {
	Content      string    `json:"content"`
	DocumentSlug string    `json:"document_slug"`
	Status       string    `json:"status"`
	Slug         string    `json:"slug"`
	UpdatedAt    time.Time `json:"updated_at"`
	InsertedAt   time.Time `json:"inserted_at"`
}
