-- name: InsertBlog :one
INSERT INTO blogs (
    owner_id, title, content, image, created_at
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
SELECT b.*, u.nickname, u.avatar FROM blogs b
JOIN users u
ON b.owner_id = u.id
WHERE b.id = $1 LIMIT 1;

-- name: ListBlogs :many
SELECT b.*, u.nickname, u.avatar FROM blogs b
JOIN users u 
ON b.owner_id = u.id
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1;

-- name: UpdateBlog :one
UPDATE blogs
SET
    title = COALESCE(sqlc.narg(title), title),
    content = COALESCE(sqlc.narg(content), content),
    image = COALESCE(sqlc.narg(image), image),
    updated_at = sqlc.arg(updated_at)
WHERE 
    id = sqlc.arg(id)
RETURNING *;

-- name: SearchBlog :many
SELECT b.*, u.nickname, u.avatar FROM blogs b 
JOIN users u ON b.owner_id = u.id
WHERE title LIKE $1
LIMIT $2
OFFSET $3;

-- name: CountBlog :one
SELECT COUNT(*) FROM blogs;