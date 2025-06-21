package db

import (
	"context"
	"fmt"
)

// DeleteRoleTxParams contains the input parameters to delete a role
type DeleteRoleTxParams struct {
	AccountID      int64   `json:"account_id"`
	ID             int64   `json:"id"`
}

// DeleteRoleTx performs the deletion of the role.
func (store *SQLStore) DeleteRoleTx(ctx context.Context, arg DeleteRoleTxParams) (error) {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Deletes Role Permissions
		err = q.DeleteRolePermissions(ctx, arg.ID)
		if err != nil {
			fmt.Println("error deleting role permissions", err)
			return err
		}

		// Deletes Role
		err = q.DeleteRole(ctx, DeleteRoleParams{
			AccountID: arg.AccountID,
			ID:        arg.ID,
		})

		if err != nil {
			fmt.Println("error deleting role", err)
			return err
		}

		return err
	})

	return err
}