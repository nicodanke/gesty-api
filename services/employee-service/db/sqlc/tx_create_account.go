package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

// CreateAccountTxParams contains the input parameters to create a account
type CreateAccountTxParams struct {
	AccountID int64 `json:"account_id"`
}

// CreateAccountTx performs the creation of the account.
func (store *SQLStore) CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		base_actions := []struct {
			Name        string
			Description string
		}{
			{"Ingreso", "Inicio de la jornada laboral"},
			{"Salida", "Finalización de la jornada laboral"},
		}

		// Creates Base Actions
		for _, action := range base_actions {
			_, err = q.CreateAction(ctx, CreateActionParams{
				AccountID:    arg.AccountID,
				Name:         action.Name,
				Description:  pgtype.Text{String: action.Description, Valid: true},
				Enabled:      true,
				CanBeDeleted: false,
			})
			if err != nil {
				fmt.Println("error creating action", err)
				return err
			}
		}

		return err
	})

	return err
}
