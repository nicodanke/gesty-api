-- name: CreateEmployeeFacility :one
INSERT INTO "employee_facility" (
    employee_id, facility_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEmployeeFacilitiesByEmployeeId :many
SELECT * FROM "employee_facility"
WHERE employee_id = $1;

-- name: GetEmployeeFacilitiesByFacilityId :many
SELECT * FROM "employee_facility"
WHERE facility_id = $1;

-- name: DeleteEmployeeFacilityByEmployeeId :exec
DELETE FROM "employee_facility"
WHERE employee_id = $1;

-- name: DeleteEmployeeFacilityByFacilityId :exec
DELETE FROM "employee_facility"
WHERE facility_id = $1;
