-- name: CreateAttendance :one
INSERT INTO "attendance" (
    id, time_in, employee_id, action_id, device_id, account_id, precision
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetAttendance :one
SELECT a.*, e.name as employee_name, d.name as device_name, ac.name as action_name
FROM "attendance" a
LEFT JOIN "employee" e ON a.employee_id = e.id
LEFT JOIN "device" d ON a.device_id = d.id
LEFT JOIN "action" ac ON a.action_id = ac.id
WHERE a.account_id = $1 AND a.id = $2 LIMIT 1;

-- name: GetAttendances :many
SELECT a.*, e.name as employee_name, d.name as device_name, ac.name as action_name
FROM "attendance" a
LEFT JOIN "employee" e ON a.employee_id = e.id
LEFT JOIN "device" d ON a.device_id = d.id
LEFT JOIN "action" ac ON a.action_id = ac.id
WHERE a.account_id = $1
ORDER BY time_in DESC
LIMIT $2
OFFSET $3;

-- name: DeleteAttendance :exec
DELETE FROM "attendance"
WHERE id = $1;
