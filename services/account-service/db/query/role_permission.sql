-- name: CreateRolePermission :one
INSERT INTO role_permission (
    role_id, permission_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetRolePermissionsByRoleId :many
SELECT * FROM role_permission
WHERE role_id = $1;

-- name: DeleteRolePermissions :exec
DELETE FROM role_permission
WHERE role_id = $1;
