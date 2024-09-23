package repo

import (
	"context"
	"github.com/rmarken5/markshop-blog/shared/models/blog"
)

type Repository interface {
	GetAll(ctx context.Context) ([]blog.Blog, error)
	GetByID(ctx context.Context, id string) (blog.Blog, error)
	GetByName(ctx context.Context, name string) (blog.Blog, error)
	CreateBlog(ctx context.Context, b blog.Blog) error
	UpdateBlog(ctx context.Context, b blog.Blog) error
	DeleteBlog(ctx context.Context, id string) error
}
