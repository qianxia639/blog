-- name: InsertArticle :one
INSERT INTO articles (
    owner_id, title, content, image, is_reward, is_critique, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: IncrViews :exec
UPDATE articles
SET
    views = views + 1
WHERE id = $1;

-- name: GetArticle :one
SELECT a.*, u.nickname, u.avatar FROM articles a
JOIN users u
ON a.owner_id = u.id
WHERE a.id = $1 LIMIT 1;

-- name: ListArticles :many
SELECT a.*, u.nickname, u.avatar FROM articles a
JOIN users u 
ON a.owner_id = u.id
WHERE title LIKE $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = $1;

-- name: UpdateArticle :one
UPDATE articles
SET
    title = COALESCE(sqlc.narg(title), title),
    content = COALESCE(sqlc.narg(content), content),
    image = COALESCE(sqlc.narg(image), image),
    updated_at = sqlc.arg(updated_at)
WHERE 
    id = sqlc.arg(id)
RETURNING *;

-- name: CountArticle :one
SELECT COUNT(*) FROM articles;