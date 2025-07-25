// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: employee_address.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEmployeeAddress = `-- name: CreateEmployeeAddress :one
INSERT INTO "employee_address" (
    employee_id, country, state, sub_state, street, number, unit, zip_code, lat, lng
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING employee_id, country, state, sub_state, street, number, unit, zip_code, lat, lng
`

type CreateEmployeeAddressParams struct {
	EmployeeID int64         `json:"employee_id"`
	Country    string        `json:"country"`
	State      string        `json:"state"`
	SubState   pgtype.Text   `json:"sub_state"`
	Street     string        `json:"street"`
	Number     string        `json:"number"`
	Unit       pgtype.Text   `json:"unit"`
	ZipCode    string        `json:"zip_code"`
	Lat        pgtype.Float8 `json:"lat"`
	Lng        pgtype.Float8 `json:"lng"`
}

func (q *Queries) CreateEmployeeAddress(ctx context.Context, arg CreateEmployeeAddressParams) (EmployeeAddress, error) {
	row := q.db.QueryRow(ctx, createEmployeeAddress,
		arg.EmployeeID,
		arg.Country,
		arg.State,
		arg.SubState,
		arg.Street,
		arg.Number,
		arg.Unit,
		arg.ZipCode,
		arg.Lat,
		arg.Lng,
	)
	var i EmployeeAddress
	err := row.Scan(
		&i.EmployeeID,
		&i.Country,
		&i.State,
		&i.SubState,
		&i.Street,
		&i.Number,
		&i.Unit,
		&i.ZipCode,
		&i.Lat,
		&i.Lng,
	)
	return i, err
}

const deleteEmployeeAddress = `-- name: DeleteEmployeeAddress :exec
DELETE FROM "employee_address"
WHERE employee_id = $1
`

func (q *Queries) DeleteEmployeeAddress(ctx context.Context, employeeID int64) error {
	_, err := q.db.Exec(ctx, deleteEmployeeAddress, employeeID)
	return err
}

const getEmployeeAddressByEmployeeId = `-- name: GetEmployeeAddressByEmployeeId :one
SELECT employee_id, country, state, sub_state, street, number, unit, zip_code, lat, lng FROM "employee_address" WHERE employee_id = $1
`

func (q *Queries) GetEmployeeAddressByEmployeeId(ctx context.Context, employeeID int64) (EmployeeAddress, error) {
	row := q.db.QueryRow(ctx, getEmployeeAddressByEmployeeId, employeeID)
	var i EmployeeAddress
	err := row.Scan(
		&i.EmployeeID,
		&i.Country,
		&i.State,
		&i.SubState,
		&i.Street,
		&i.Number,
		&i.Unit,
		&i.ZipCode,
		&i.Lat,
		&i.Lng,
	)
	return i, err
}

const updateEmployeeAddress = `-- name: UpdateEmployeeAddress :one
UPDATE "employee_address"
SET
    country = COALESCE($1, country),
    state = COALESCE($2, state),
    sub_state = COALESCE($3, sub_state),
    street = COALESCE($4, street),
    number = COALESCE($5, number),
    unit = COALESCE($6, unit),
    zip_code = COALESCE($7, zip_code),
    lat = COALESCE($8, lat),
    lng = COALESCE($9, lng)
WHERE
    employee_id = $10
RETURNING employee_id, country, state, sub_state, street, number, unit, zip_code, lat, lng
`

type UpdateEmployeeAddressParams struct {
	Country    pgtype.Text   `json:"country"`
	State      pgtype.Text   `json:"state"`
	SubState   pgtype.Text   `json:"sub_state"`
	Street     pgtype.Text   `json:"street"`
	Number     pgtype.Text   `json:"number"`
	Unit       pgtype.Text   `json:"unit"`
	ZipCode    pgtype.Text   `json:"zip_code"`
	Lat        pgtype.Float8 `json:"lat"`
	Lng        pgtype.Float8 `json:"lng"`
	EmployeeID int64         `json:"employee_id"`
}

func (q *Queries) UpdateEmployeeAddress(ctx context.Context, arg UpdateEmployeeAddressParams) (EmployeeAddress, error) {
	row := q.db.QueryRow(ctx, updateEmployeeAddress,
		arg.Country,
		arg.State,
		arg.SubState,
		arg.Street,
		arg.Number,
		arg.Unit,
		arg.ZipCode,
		arg.Lat,
		arg.Lng,
		arg.EmployeeID,
	)
	var i EmployeeAddress
	err := row.Scan(
		&i.EmployeeID,
		&i.Country,
		&i.State,
		&i.SubState,
		&i.Street,
		&i.Number,
		&i.Unit,
		&i.ZipCode,
		&i.Lat,
		&i.Lng,
	)
	return i, err
}
