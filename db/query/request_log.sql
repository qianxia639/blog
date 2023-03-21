-- name: InsertRequestLog :one
INSERT INTO request_logs (
    method, path, status_code, ip, hostname, request_body, response_time, content_type, user_agent
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;