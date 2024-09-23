package postgres

import (
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/rmarken5/markshop-blog/shared/models/blog"
	sharedtime "github.com/rmarken5/markshop-blog/shared/time"
	"testing"
)

func TestPostgres_GetAll(t *testing.T) {
	var (
		haveBlog = blog.Blog{
			ID:        uuid.New(),
			Content:   "content",
			CreatedAt: sharedtime.TimeNow(),
			CreatedBy: "ryan",
			UpdatedAt: sharedtime.TimeNow(),
			UpdatedBy: "marken",
			ThumbNail: "derp",
			TagNames:  []string{},
		}
	)
	testCases := map[string]struct {
		mockDb func(mock pgxmock.PgxPoolIface)
		wantBlogs []blog.Blog
	} {
		"should return active blogs" :{
			mockDb: func(mock pgxmock.PgxPoolIface) {
				rows := pgxmock.NewRows([]string{uuid.NewString(), }),
				mock.ExpectQuery(getAllActiveBlogs,).WillReturnRows(),
			}
		},
	}

	mock, err := pgxmock.NewPool()
}
