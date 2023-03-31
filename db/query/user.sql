-- name: CreateUser :one
INSERT INTO users (
    username, email, nickname, password, register_time
) VALUES(
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1
OR email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
    nickname = COALESCE(sqlc.narg(nickname), nickname),
    email = COALESCE(sqlc.narg(email), email),
    password = COALESCE(sqlc.narg(password), password),
    avatar = COALESCE(sqlc.narg(avatar), avatar)
WHERE
    username = sqlc.arg(username)
RETURNING *;