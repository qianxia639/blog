// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: blog.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const countBlog = `-- name: CountBlog :one
SELECT COUNT(*) FROM blogs
`

func (q *Queries) CountBlog(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countBlog)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1
`

func (q *Queries) DeleteBlog(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBlog, id)
	return err
}

const getBlog = `-- name: GetBlog :one
SELECT b.id, b.owner_id, b.title, b.content, b.image, b.views, b.created_at, b.updated_at, u.nickname, u.avatar FROM blogs b
JOIN users u
ON b.owner_id = u.id
WHERE b.id = $1 LIMIT 1
`

type GetBlogRow struct {
	ID        int64     `json:"id"`
	OwnerID   int64     `json:"owner_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	Views     int32     `json:"views"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
}

func (q *Queries) GetBlog(ctx context.Context, id int64) (GetBlogRow, error) {
	row := q.db.QueryRowContext(ctx, getBlog, id)
	var i GetBlogRow
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Title,
		&i.Content,
		&i.Image,
		&i.Views,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Nickname,
		&i.Avatar,
	)
	return i, err
}

const incrViews = `-- name: IncrViews :exec
UPDATE blogs
SET
    views = views + 1
WHERE id = $1
`

func (q *Queries) IncrViews(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, incrViews, id)
	return err
}

const insertBlog = `-- name: InsertBlog :one
INSERT INTO blogs (
    owner_id, title, content, image, created_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, owner_id, title, content, image, views, created_at, updated_at
`

type InsertBlogParams struct {
	OwnerID   int64     `json:"owner_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) InsertBlog(ctx context.Context, arg InsertBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, insertBlog,
		arg.OwnerID,
		arg.Title,
		arg.Content,
		arg.Image,
		arg.CreatedAt,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Title,
		&i.Content,
		&i.Image,
		&i.Views,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listBlogs = `-- name: ListBlogs :many
SELECT b.id, b.owner_id, b.title, b.content, b.image, b.views, b.created_at, b.updated_at, u.nickname, u.avatar FROM blogs b
JOIN users u 
ON b.owner_id = u.id
ORDER BY created_at
LIMIT $1
OFFSET $2
`

type ListBlogsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListBlogsRow struct {
	ID        int64     `json:"id"`
	OwnerID   int64     `json:"owner_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	Views     int32     `json:"views"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
}

func (q *Queries) ListBlogs(ctx context.Context, arg ListBlogsParams) ([]ListBlogsRow, error) {
	rows, err := q.db.QueryContext(ctx, listBlogs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListBlogsRow{}
	for rows.Next() {
		var i ListBlogsRow
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Title,
			&i.Content,
			&i.Image,
			&i.Views,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Nickname,
			&i.Avatar,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchBlog = `-- name: SearchBlog :many
SELECT b.id, b.owner_id, b.title, b.content, b.image, b.views, b.created_at, b.updated_at, u.nickname, u.avatar FROM blogs b 
JOIN users u ON b.owner_id = u.id
WHERE title LIKE $1
LIMIT $2
OFFSET $3
`

type SearchBlogParams struct {
	Title  string `json:"title"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type SearchBlogRow struct {
	ID        int64     `json:"id"`
	OwnerID   int64     `json:"owner_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	Views     int32     `json:"views"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
}

func (q *Queries) SearchBlog(ctx context.Context, arg SearchBlogParams) ([]SearchBlogRow, error) {
	rows, err := q.db.QueryContext(ctx, searchBlog, arg.Title, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchBlogRow{}
	for rows.Next() {
		var i SearchBlogRow
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Title,
			&i.Content,
			&i.Image,
			&i.Views,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Nickname,
			&i.Avatar,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBlog = `-- name: UpdateBlog :one
UPDATE blogs
SET
    title = COALESCE($1, title),
    content = COALESCE($2, content),
    image = COALESCE($3, image),
    updated_at = $4
WHERE 
    id = $5
RETURNING id, owner_id, title, content, image, views, created_at, updated_at
`

type UpdateBlogParams struct {
	Title     sql.NullString `json:"title"`
	Content   sql.NullString `json:"content"`
	Image     sql.NullString `json:"image"`
	UpdatedAt time.Time      `json:"updated_at"`
	ID        int64          `json:"id"`
}

func (q *Queries) UpdateBlog(ctx context.Context, arg UpdateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, updateBlog,
		arg.Title,
		arg.Content,
		arg.Image,
		arg.UpdatedAt,
		arg.ID,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Title,
		&i.Content,
		&i.Image,
		&i.Views,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
