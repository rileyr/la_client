package littleaspen

import "time"

type Document struct {
	Slug       string    `json:"slug"`
	Title      string    `json:"title"`
	UserID     string    `json:"user_id"`
	InsertedAt time.Time `json:"inserted_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
