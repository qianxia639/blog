-- name: InsertType :one
INSERT INTO types (
    type_name
) VALUES (
    $1
)
RETURNING *;