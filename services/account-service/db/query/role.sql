-- name: CreateRole :one
INSERT INTO "role" (
    account_id, name, description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetRole :one
SELECT r.id, r.name, r.description,
    COALESCE(
        ARRAY_AGG(rp.permission_id) FILTER (WHERE rp.permission_id IS NOT NULL),
        '{}'::int8[]
    ) as permission_ids
FROM "role" r
LEFT JOIN role_permission rp ON r.id = rp.role_id
WHERE r.account_id = $1 AND id = $2
GROUP BY r.id, r.name, r.description
LIMIT 1;

-- name: GetRoles :many
SELECT r.id, r.name, r.description,
    COALESCE(
        ARRAY_AGG(rp.permission_id) FILTER (WHERE rp.permission_id IS NOT NULL),
        '{}'::int8[]
    ) as permission_ids
FROM "role" r
LEFT JOIN role_permission rp ON r.id = rp.role_id
WHERE r.account_id = $1
GROUP BY r.id, r.name, r.description
ORDER BY LOWER(r.name)
LIMIT $2
OFFSET $3;

-- name: DeleteRole :exec
DELETE FROM "role"
WHERE account_id = $1 AND id = $2;

-- name: UpdateRole :one
UPDATE "role"
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at)
WHERE
    account_id = sqlc.arg(account_id) AND id = sqlc.arg(id)
RETURNING *;
