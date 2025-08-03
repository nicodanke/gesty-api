-- name: CreateAttendance :one
INSERT INTO "attendance" (
    time_in, employee_id, action_id, device_id, precision
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetAttendance :one
SELECT * FROM "attendance"
WHERE id = $1 LIMIT 1;

-- name: GetAttendances :many
SELECT * FROM "attendance"
ORDER BY time_in DESC
LIMIT $1
OFFSET $2;

-- name: DeleteAttendance :exec
DELETE FROM "attendance"
WHERE id = $1;
