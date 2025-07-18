package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

// CreateRoleTxParams contains the input parameters to create a role
type CreateRoleTxParams struct {
	AccountID     int64   `json:"account_id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	PermissionIDs []int64 `json:"permission_ids"`
}

// CreateRoleTxResult is the result of the role creation
type CreateRoleTxResult struct {
	Role          Role    `json:"role"`
	PermissionIDs []int64 `json:"permissions_ids"`
}

// CreateRoleTx performs the creation of the role and the permissions.
func (store *SQLStore) CreateRoleTx(ctx context.Context, arg CreateRoleTxParams) (CreateRoleTxResult, error) {
	var result CreateRoleTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Creates Role
		result.Role, err = q.CreateRole(ctx, CreateRoleParams{
			AccountID: arg.AccountID,
			Name:      arg.Name,
			Description: pgtype.Text{
				String: arg.Description,
				Valid:  arg.Description != "",
			},
		})

		if err != nil {
			fmt.Println("error creating role", err)
			return err
		}

		// Assign all permissions to admin role
		for _, value := range arg.PermissionIDs {
			_, err = q.CreateRolePermission(ctx, CreateRolePermissionParams{
				RoleID:       result.Role.ID,
				PermissionID: int64(value),
			})
			if err != nil {
				return err
			}
		}

		result.PermissionIDs = arg.PermissionIDs

		return err
	})

	return result, err
}
