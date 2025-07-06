-- name: GetPermissions :many
SELECT * FROM "permission"
LIMIT $1
OFFSET $2;
