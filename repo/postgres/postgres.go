package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rmarken5/markshop-blog/repo"
	"github.com/rmarken5/markshop-blog/shared/models/blog"
)

var _ repo.Repository = Postgres{}

type (
	Postgres struct {
		db *sqlx.DB
	}
)

func New(db *sqlx.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

// language=postgresql
const getAllActiveBlogs = `select p.*, array_agg(t.name) AS tag_names
from post p
         inner join public.post_tag pt on p.id = pt.post_id
         inner join tag t on pt.tag_id = t.id
where p.deleted_at is null
GROUP BY p.id`

func (p Postgres) GetAll(ctx context.Context) ([]blog.Blog, error) {
	var blogs []blog.Blog
	err := p.db.SelectContext(ctx, &blogs, getAllActiveBlogs)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Join(err, repo.ErrNotFound)
		default:
			return nil, errors.Join(err, repo.ErrGeneral)
		}
	}

	return blogs, nil
}

func (p Postgres) GetByID(ctx context.Context, id string) (blog.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) GetByName(ctx context.Context, name string) (blog.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) CreateBlog(ctx context.Context, b blog.Blog) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) UpdateBlog(ctx context.Context, b blog.Blog) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) DeleteBlog(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
