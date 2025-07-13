-- name: CreateFacility :one
INSERT INTO "facility" (
    account_id, name, description, open_time, close_time
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetFacility :one
SELECT f.id, f.name, f.description, f.open_time, f.close_time, fa.country, fa.state, fa.sub_state, fa.street, fa.number, fa.unit, fa.postal_code, fa.lat, fa.lng
FROM "facility" f
LEFT JOIN facility_address fa ON f.id = fa.facility_id
WHERE f.account_id = $1 AND f.id = $2
GROUP BY f.id, f.name, f.description, f.open_time, f.close_time, fa.country, fa.state, fa.sub_state, fa.street, fa.number, fa.unit, fa.postal_code, fa.lat, fa.lng
LIMIT 1;

-- name: GetFacilities :many
SELECT f.id, f.name, f.description, f.open_time, f.close_time, fa.country, fa.state, fa.sub_state, fa.street, fa.number, fa.unit, fa.postal_code, fa.lat, fa.lng
FROM "facility" f
LEFT JOIN facility_address fa ON f.id = fa.facility_id
WHERE f.account_id = $1
GROUP BY f.id, f.name, f.description, f.open_time, f.close_time, fa.country, fa.state, fa.sub_state, fa.street, fa.number, fa.unit, fa.postal_code, fa.lat, fa.lng
ORDER BY LOWER(f.name)
LIMIT $2
OFFSET $3;

-- name: DeleteFacility :exec
DELETE FROM "facility"
WHERE account_id = $1 AND id = $2;

-- name: UpdateFacility :one
UPDATE "facility"
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    open_time = COALESCE(sqlc.narg(open_time), open_time),
    close_time = COALESCE(sqlc.narg(close_time), close_time),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at)
WHERE
    account_id = sqlc.arg(account_id) AND id = sqlc.arg(id)
RETURNING *;
