-- name: CreateEmployeePhoto :one
INSERT INTO "employee_photo" (
    employee_id, image_base_64, vector_image, account_id, is_profile
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetEmployeePhotos :many
SELECT * FROM "employee_photo"
WHERE employee_id = $1 AND account_id = $2
ORDER BY created_at DESC
LIMIT $3
OFFSET $4;

-- name: GetEmployeeProfilePhoto :one
SELECT * FROM "employee_photo"
WHERE employee_id = $1 AND account_id = $2 AND is_profile = true
ORDER BY created_at DESC
LIMIT 3;

-- name: DeleteEmployeePhoto :exec
DELETE FROM "employee_photo"
WHERE id = $1 AND employee_id = $2 AND is_profile = false;
