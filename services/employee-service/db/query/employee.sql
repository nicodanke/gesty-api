-- name: CreateEmployee :one
INSERT INTO "employee" (
    account_id, name, lastname, email, phone, gender, real_id, fiscal_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetEmployee :one
SELECT e.id, e.name, e.lastname, e.email, e.phone, e.gender, e.real_id, e.fiscal_id, ea.country, ea.state, ea.sub_state, ea.street, ea.number, ea.unit, ea.zip_code, ea.lat, ea.lng,
    COALESCE(
        ARRAY_AGG(ef.facility_id) FILTER (WHERE ef.facility_id IS NOT NULL),
        '{}'::int8[]
    ) as facility_ids
FROM "employee" e
LEFT JOIN employee_address ea ON e.id = ea.employee_id
LEFT JOIN employee_facility ef ON e.id = ef.employee_id
WHERE e.account_id = $1 AND e.id = $2
GROUP BY e.id, e.name, e.lastname, e.email, e.phone, e.gender, e.real_id, e.fiscal_id, ea.country, ea.state, ea.sub_state, ea.street, ea.number, ea.unit, ea.zip_code, ea.lat, ea.lng
LIMIT 1;

-- name: GetEmployees :many
SELECT e.id, e.name, e.lastname, e.email, e.phone, e.gender, e.real_id, e.fiscal_id, ea.country, ea.state, ea.sub_state, ea.street, ea.number, ea.unit, ea.zip_code, ea.lat, ea.lng,
    COALESCE(
        ARRAY_AGG(ef.facility_id) FILTER (WHERE ef.facility_id IS NOT NULL),
        '{}'::int8[]
    ) as facility_ids
FROM "employee" e
LEFT JOIN employee_address ea ON e.id = ea.employee_id
LEFT JOIN employee_facility ef ON e.id = ef.employee_id
WHERE e.account_id = $1
GROUP BY e.id, e.name, e.lastname, e.email, e.phone, e.gender, e.real_id, e.fiscal_id, ea.country, ea.state, ea.sub_state, ea.street, ea.number, ea.unit, ea.zip_code, ea.lat, ea.lng
ORDER BY LOWER(e.name)
LIMIT $2
OFFSET $3;

-- name: DeleteEmployee :exec
DELETE FROM "employee"
WHERE account_id = $1 AND id = $2;

-- name: UpdateEmployee :one
UPDATE "employee"
SET
    name = COALESCE(sqlc.narg(name), name),
    lastname = COALESCE(sqlc.narg(lastname), lastname),
    email = COALESCE(sqlc.narg(email), email),
    phone = COALESCE(sqlc.narg(phone), phone),
    gender = COALESCE(sqlc.narg(gender), gender),
    real_id = COALESCE(sqlc.narg(real_id), real_id),
    fiscal_id = COALESCE(sqlc.narg(fiscal_id), fiscal_id),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at)
WHERE
    account_id = sqlc.arg(account_id) AND id = sqlc.arg(id)
RETURNING *;
