-- name: CreateUser :one
INSERT INTO "user" (
    account_id, name, lastname, email, username, password, phone, is_admin, role_id, active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE account_id = $1 AND id = $2 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM "user"
WHERE username = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM "user"
WHERE account_id = $1
ORDER BY LOWER(name), LOWER(lastname)
LIMIT $2
OFFSET $3;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE account_id = $1 AND id = $2;

-- name: GetUserPermissions :many
SELECT p.code FROM "user" u
JOIN "role" r ON r.id = u.role_id
JOIN "role_permission" rp ON rp.role_id = r.id
JOIN "permission" p ON p.id = rp.permission_id
WHERE u.id = $1;

-- name: UpdateUser :one
UPDATE "user"
SET
    name = COALESCE(sqlc.narg(name), name),
    lastname = COALESCE(sqlc.narg(lastname), lastname),
    email = COALESCE(sqlc.narg(email), email),
    role_id = COALESCE(sqlc.narg(role_id), role_id),
    phone = COALESCE(sqlc.narg(phone), phone),
    active = COALESCE(sqlc.narg(active), active),
    is_admin = COALESCE(sqlc.narg(is_admin), is_admin),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at)
WHERE
    account_id = sqlc.arg(account_id) AND id = sqlc.arg(id)
RETURNING *;
