-- name: CreateRole :one
INSERT INTO "role" (
    account_id, name
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetRole :one
SELECT * FROM "role"
WHERE account_id = $1 AND id = $2 LIMIT 1;

-- name: GetRoles :many
SELECT * FROM "role"
WHERE account_id = $1
ORDER BY name
LIMIT $2
OFFSET $3;

-- name: DeleteRole :exec
DELETE FROM "role"
WHERE account_id = $1 AND id = $2;
