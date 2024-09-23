package blog

import (
	"github.com/google/uuid"
	"time"
)

type (
	Tag struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Blogs     []Blog    `json:"blogs"`
		CreatedAt time.Time `json:"created_at"`
	}
)
