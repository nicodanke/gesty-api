package db

import (
	"context"
	"fmt"
)

// DeleteFacilityTxParams contains the input parameters to delete a facility
type DeleteFacilityTxParams struct {
	AccountID int64 `json:"account_id"`
	ID        int64 `json:"id"`
}

// DeleteFacilityTx performs the deletion of the facility.
func (store *SQLStore) DeleteFacilityTx(ctx context.Context, arg DeleteFacilityTxParams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Deletes Facility Address
		err = q.DeleteFacilityAddress(ctx, arg.ID)
		if err != nil {
			fmt.Println("error deleting facility address", err)
			return err
		}

		// Deletes Employee Facility
		err = q.DeleteEmployeeFacilityByFacilityId(ctx, arg.ID)
		if err != nil {
			fmt.Println("error deleting employee-facility", err)
			return err
		}

		// Deletes Facility
		err = q.DeleteFacility(ctx, DeleteFacilityParams(arg))

		if err != nil {
			fmt.Println("error deleting facility", err)
			return err
		}

		return err
	})

	return err
}
