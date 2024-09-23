package blog

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

type (
	Blog struct {
		ID          uuid.UUID      `db:"id"`
		Title       string         `db:"title"`
		Description string         `db:"description"`
		Content     string         `db:"content"`
		CreatedAt   time.Time      `db:"created_at"`
		CreatedBy   string         `db:"created_by"`
		UpdatedAt   time.Time      `db:"updated_at"`
		UpdatedBy   string         `db:"updated_by"`
		ThumbNail   string         `db:"thumbnail"`
		TagNames    pq.StringArray `db:"tag_names"`
		DeletedAt   *time.Time     `db:"deleted_at"`
	}
)
