-- name: CreateDeviceAction :one
INSERT INTO "device_action" (
    device_id, action_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetDeviceActionsByDeviceId :many
SELECT * FROM "device_action"
WHERE device_id = $1;

-- name: DeleteDeviceActionByDeviceId :exec
DELETE FROM "device_action"
WHERE device_id = $1;
