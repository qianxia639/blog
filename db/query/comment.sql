-- name: InsertComment :one
INSERT INTO comments (
    blog_id, comment_id, nickname, avatar, content
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;