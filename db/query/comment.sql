-- name: CreateComment :one
INSERT INTO comments (
    owner_id, parent_id, nickname, avatar, content
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetComments :many
SELECT * FROM comments
WHERE owner_id = $1;

-- name: GetChildComments :many
SELECT * FROM comments
WHERE id = $1 AND parent_id = $2;