package littleaspen

import "time"

type Acceptance struct {
	ContentVersionSlug string                 `json:"content_version_slug,omitempty"`
	DocumentSlug       string                 `json:"document_slug,omitempty"`
	ExternalID         string                 `json:"id"`
	InsertedAt         *time.Time             `json:"inserted_at,omitempty"`
	UpdatedAt          *time.Time             `json:"updated_at,omitempty"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
}
