-- name: CreateEmployeeAddress :one
INSERT INTO "employee_address" (
    employee_id, country, state, sub_state, street, number, unit, zip_code, lat, lng
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetEmployeeAddressByEmployeeId :one
SELECT * FROM "employee_address" WHERE employee_id = $1;

-- name: DeleteEmployeeAddress :exec
DELETE FROM "employee_address"
WHERE employee_id = $1;

-- name: UpdateEmployeeAddress :one
UPDATE "employee_address"
SET
    country = COALESCE(sqlc.narg(country), country),
    state = COALESCE(sqlc.narg(state), state),
    sub_state = COALESCE(sqlc.narg(sub_state), sub_state),
    street = COALESCE(sqlc.narg(street), street),
    number = COALESCE(sqlc.narg(number), number),
    unit = COALESCE(sqlc.narg(unit), unit),
    zip_code = COALESCE(sqlc.narg(zip_code), zip_code),
    lat = COALESCE(sqlc.narg(lat), lat),
    lng = COALESCE(sqlc.narg(lng), lng)
WHERE
    employee_id = sqlc.arg(employee_id)
RETURNING *;
