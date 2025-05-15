-- name: CreateRolePermission :one
INSERT INTO role_permission (
    role_id, permission_id
) VALUES (
    $1, $2
) RETURNING *;
