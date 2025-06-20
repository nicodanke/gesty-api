-- name: CreateAccount :one
INSERT INTO account (
    code, company_name, email, active
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: GetAccountByCode :one
SELECT * FROM account
WHERE code = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY company_name
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE account
SET
    company_name = COALESCE(sqlc.narg(company_name), company_name),
    phone = COALESCE(sqlc.narg(phone), phone),
    email = COALESCE(sqlc.narg(email), email),
    web_url = COALESCE(sqlc.narg(web_url), web_url),
    active = COALESCE(sqlc.narg(active), active),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at)
WHERE
    id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;
