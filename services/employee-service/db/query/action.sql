-- name: CreateAction :one
INSERT INTO "action" (
    account_id, name, description, enabled, can_be_deleted
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetAction :one
SELECT * FROM "action"
WHERE account_id = $1 AND id = $2 LIMIT 1;

-- name: GetActions :many
SELECT * FROM "action"
WHERE account_id = $1
ORDER BY LOWER(name)
LIMIT $2
OFFSET $3;

-- name: DeleteAction :exec
DELETE FROM "action"
WHERE account_id = $1 AND id = $2;

-- name: UpdateAction :one
UPDATE "action"
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    enabled = COALESCE(sqlc.narg(enabled), enabled),
    can_be_deleted = COALESCE(sqlc.narg(can_be_deleted), can_be_deleted)
WHERE
    account_id = sqlc.arg(account_id) AND id = sqlc.arg(id)
RETURNING *;
