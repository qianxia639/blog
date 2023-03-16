-- name: LoginUser :one
SELECT * FROM users
WHERE (username = $1 AND password = $2)
OR (email = $1 AND password = $2) LIMIT 1;