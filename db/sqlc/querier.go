// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
)

type Querier interface {
	CountBlog(ctx context.Context) (int64, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteBlog(ctx context.Context, id int64) error
	GetBlog(ctx context.Context, id int64) (GetBlogRow, error)
	GetChildComments(ctx context.Context, arg GetChildCommentsParams) ([]Comment, error)
	GetComments(ctx context.Context, ownerID int64) ([]Comment, error)
	GetUser(ctx context.Context, username string) (User, error)
	IncrViews(ctx context.Context, id int64) error
	InsertBlog(ctx context.Context, arg InsertBlogParams) (Blog, error)
	InsertRequestLog(ctx context.Context, arg InsertRequestLogParams) (RequestLog, error)
	ListBlogs(ctx context.Context, arg ListBlogsParams) ([]ListBlogsRow, error)
	SearchBlog(ctx context.Context, arg SearchBlogParams) ([]SearchBlogRow, error)
	UpdateBlog(ctx context.Context, arg UpdateBlogParams) (Blog, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
