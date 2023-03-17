-- name: InsertBlog :one
INSERT INTO blogs (
    owner_id, type_id, title, content, image
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: IncrViews :exec
UPDATE blogs
SET
    views = views + 1
WHERE id = $1;