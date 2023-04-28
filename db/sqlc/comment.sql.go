// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: comment.sql

package db

import (
	"context"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (
    owner_id, parent_id, nickname, avatar, content
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, owner_id, parent_id, nickname, avatar, content, created_at
`

type CreateCommentParams struct {
	OwnerID  int64  `json:"owner_id"`
	ParentID int64  `json:"parent_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Content  string `json:"content"`
}

func (q *Queries) CreateComment(ctx context.Context, arg *CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment,
		arg.OwnerID,
		arg.ParentID,
		arg.Nickname,
		arg.Avatar,
		arg.Content,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.ParentID,
		&i.Nickname,
		&i.Avatar,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const getChildComments = `-- name: GetChildComments :many
SELECT id, owner_id, parent_id, nickname, avatar, content, created_at FROM comments
WHERE id = $1 AND parent_id = $2
`

type GetChildCommentsParams struct {
	ID       int64 `json:"id"`
	ParentID int64 `json:"parent_id"`
}

func (q *Queries) GetChildComments(ctx context.Context, arg *GetChildCommentsParams) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, getChildComments, arg.ID, arg.ParentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.ParentID,
			&i.Nickname,
			&i.Avatar,
			&i.Content,
			&i.CreatedAt,
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

const getComments = `-- name: GetComments :many
SELECT id, owner_id, parent_id, nickname, avatar, content, created_at FROM comments
WHERE owner_id = $1
`

func (q *Queries) GetComments(ctx context.Context, ownerID int64) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, getComments, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.ParentID,
			&i.Nickname,
			&i.Avatar,
			&i.Content,
			&i.CreatedAt,
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
