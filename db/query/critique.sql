-- name: CreateCritique :one
INSERT INTO critiques (
    owner_id, parent_id, nickname, avatar, content
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetCritiques :many
SELECT * FROM critiques
WHERE owner_id = $1 AND parent_id = 0;

-- name: GetChildCritiques :many
SELECT * FROM critiques
WHERE parent_id = $1;