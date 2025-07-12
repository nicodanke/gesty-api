-- name: CreateDeviceHealth :one
INSERT INTO "device_health" (
    connection_type, free_memory, free_storage, battery_level, battery_save_mode, device_id
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetDeviceHealth :one
SELECT * FROM "device_health"
WHERE id = $1 LIMIT 1;

-- name: GetDeviceHealths :many
SELECT * FROM "device_health"
WHERE device_id = $1
ORDER BY created_at DESC;

-- name: DeleteDeviceHealth :exec
DELETE FROM "device_health"
WHERE device_id = $1;
