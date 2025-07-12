-- name: CreateDevice :one
INSERT INTO "device" (
    account_id, name, enabled, facility_id
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetDevice :one
SELECT d.*,
    COALESCE(
        ARRAY_AGG(da.action_id) FILTER (WHERE da.action_id IS NOT NULL),
        '{}'::int8[]
    ) as action_ids
FROM "device" d
LEFT JOIN device_action da ON d.id = da.device_id
WHERE d.account_id = $1 AND d.id = $2
GROUP BY d.id, d.name, d.enabled, d.facility_id
LIMIT 1;

-- name: GetDevices :many
SELECT * FROM "device"
WHERE account_id = $1
ORDER BY LOWER(name)
LIMIT $2
OFFSET $3;

-- name: DeleteDevice :exec
DELETE FROM "device"
WHERE account_id = $1 AND id = $2;

-- name: UpdateDevice :one
UPDATE "device"
SET
    name = COALESCE(sqlc.narg(name), name),
    password = COALESCE(sqlc.narg(password), password),
    enabled = COALESCE(sqlc.narg(enabled), enabled),
    active = COALESCE(sqlc.narg(active), active),
    activation_code = COALESCE(sqlc.narg(activation_code), activation_code),
    activation_code_expires_at = COALESCE(sqlc.narg(activation_code_expires_at), activation_code_expires_at),
    device_name = COALESCE(sqlc.narg(device_name), device_name),
    device_brand = COALESCE(sqlc.narg(device_brand), device_brand),
    device_model = COALESCE(sqlc.narg(device_model), device_model),
    device_serial_number = COALESCE(sqlc.narg(device_serial_number), device_serial_number),
    device_os = COALESCE(sqlc.narg(device_os), device_os),
    device_ram = COALESCE(sqlc.narg(device_ram), device_ram),
    device_storage = COALESCE(sqlc.narg(device_storage), device_storage),
    device_os_version = COALESCE(sqlc.narg(device_os_version), device_os_version),
    facility_id = COALESCE(sqlc.narg(facility_id), facility_id),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at)
WHERE
    account_id = sqlc.arg(account_id) AND id = sqlc.arg(id)
RETURNING *;
