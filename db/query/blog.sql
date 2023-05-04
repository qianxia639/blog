-- name: InsertArticle :one
INSERT INTO blogs (
    owner_id, title, content, image, is_reward, is_critique, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: IncrViews :exec
UPDATE blogs
SET
    views = views + 1
WHERE id = $1;

-- name: GetArticle :one
SELECT b.*, u.nickname, u.avatar FROM blogs b
JOIN users u
ON b.owner_id = u.id
WHERE b.id = $1 LIMIT 1;

-- name: ListArticles :many
SELECT b.*, u.nickname, u.avatar FROM blogs b
JOIN users u 
ON b.owner_id = u.id
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: DeleteArticle :exec
DELETE FROM blogs
WHERE id = $1;

-- name: UpdateArticle :one
UPDATE blogs
SET
    title = COALESCE(sqlc.narg(title), title),
    content = COALESCE(sqlc.narg(content), content),
    image = COALESCE(sqlc.narg(image), image),
    updated_at = sqlc.arg(updated_at)
WHERE 
    id = sqlc.arg(id)
RETURNING *;

-- name: SearchArticle :many
SELECT b.*, u.nickname, u.avatar FROM blogs b 
JOIN users u ON b.owner_id = u.id
WHERE title LIKE $1
LIMIT $2
OFFSET $3;

-- name: CountArticle :one
SELECT COUNT(*) FROM blogs;