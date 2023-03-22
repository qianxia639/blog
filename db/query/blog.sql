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

-- name: GetBlog :one
SELECT * FROM blogs
WHERE id = $1 LIMIT 1;

-- name: ListBlogs :many
SELECT * FROM blogs
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1;

-- name: UpdateBlog :one
UPDATE blogs
SET
    type_id = COALESCE(sqlc.narg(type_id), type_id),
    title = COALESCE(sqlc.narg(title), title),
    content = COALESCE(sqlc.narg(content), content),
    image = COALESCE(sqlc.narg(image), image)
WHERE 
    id = sqlc.arg(id)
RETURNING *;

-- name: SearchBlog :many
SELECT * FROM blogs
WHERE title LIKE $1;