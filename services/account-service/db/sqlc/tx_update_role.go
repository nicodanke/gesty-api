package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// UpdateRoleTxParams contains the input parameters to update a role
type UpdateRoleTxParams struct {
	AccountID     int64       `json:"account_id"`
	ID            int64       `json:"id"`
	Name          pgtype.Text `json:"name"`
	Description   pgtype.Text `json:"description"`
	PermissionIDs []int64     `json:"permission_ids"`
}

// UpdateRoleTxResult is the result of the role update
type UpdateRoleTxResult struct {
	Role          Role    `json:"role"`
	PermissionIDs []int64 `json:"permissions_ids"`
}

// UpdateRoleTx performs the update of the role and the permissions.
func (store *SQLStore) UpdateRoleTx(ctx context.Context, arg UpdateRoleTxParams) (UpdateRoleTxResult, error) {
	var result UpdateRoleTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Updates Role
		result.Role, err = q.UpdateRole(ctx, UpdateRoleParams{
			AccountID:   arg.AccountID,
			ID:          arg.ID,
			Name:        arg.Name,
			Description: arg.Description,
			UpdatedAt: pgtype.Timestamptz{
				Time:  time.Now().UTC(),
				Valid: true,
			},
		})

		if err != nil {
			fmt.Println("error updating role", err)
			return err
		}

		if arg.PermissionIDs != nil {
			err = q.DeleteRolePermissions(ctx, arg.ID)
			if err != nil {
				fmt.Println("error deleting role permissions", err)
				return err
			}

			// Assign all permissions to role
			for _, value := range arg.PermissionIDs {
				_, err = q.CreateRolePermission(ctx, CreateRolePermissionParams{
					RoleID:       result.Role.ID,
					PermissionID: value,
				})
				if err != nil {
					fmt.Println("error creating role permissions", err)
					return err
				}
			}

			result.PermissionIDs = arg.PermissionIDs
		} else {
			permissions, err := q.GetRolePermissionsByRoleId(ctx, arg.ID)
			if err != nil {
				fmt.Println("error getting role permissions", err)
				return err
			}

			permissionIDs := make([]int64, len(permissions))
			for i, permission := range permissions {
				permissionIDs[i] = permission.PermissionID
			}

			result.PermissionIDs = permissionIDs
		}

		return err
	})

	return result, err
}
