package db

import (
	"context"
	"fmt"
)

// CreateAccountTxParams contains the input parameters to create an account
type CreateAccountTxParams struct {
	Code           string `json:"code"`
	CompanyName    string `json:"company_name"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Lastname       string `json:"lastname"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
}

// CreateAccountTxResult is the result of the account creation
type CreateAccountTxResult struct {
	Account Account `json:"account"`
	User    User    `json:"user"`
}

// CreateAccountTx performs the creation of the account and the first user and also, the base entities of the account.
func (store *SQLStore) CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (CreateAccountTxResult, error) {
	var result CreateAccountTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		fmt.Println("arg.Code", arg.Code)
		// Creates account
		result.Account, err = q.CreateAccount(ctx, CreateAccountParams{
			Code:        arg.Code,
			CompanyName: arg.CompanyName,
			Email:       arg.Email,
			Active:      true,
		})
		if err != nil {
			fmt.Println("error creating account", err)
			return err
		}

		fmt.Println("result.Account", result.Account)

		// Creates Admin Role
		role, err := q.CreateRole(ctx, CreateRoleParams{
			AccountID: result.Account.ID,
			Name:      "Admin",
		})
		if err != nil {
			fmt.Println("error creating role", err)
			return err
		}

		fmt.Println("role", role)

		// Assign all permissions to admin role
		permissionIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

		for _, value := range permissionIDs {
			_, err = q.CreateRolePermission(ctx, CreateRolePermissionParams{
				RoleID:       role.ID,
				PermissionID: int64(value),
			})
			if err != nil {
				return err
			}
		}

		fmt.Println("role", role)

		// Creates user
		result.User, err = q.CreateUser(ctx, CreateUserParams{
			AccountID: result.Account.ID,
			Name:      arg.Name,
			Lastname:  arg.Lastname,
			Email:     arg.Email,
			Username:  arg.Username + "@" + arg.Code,
			Password:  arg.HashedPassword,
			RoleID:    role.ID,
			Active:    true,
			IsAdmin:   true,
		})
		if err != nil {
			return err
		}

		fmt.Println("result.User", result.User)

		return err
	})

	return result, err
}