package db

import (
	"context"
	"fmt"
)

// DeleteEmployeeParams contains the input parameters to delete an employee
type DeleteEmployeeTxParams struct {
	AccountID int64 `json:"account_id"`
	ID        int64 `json:"id"`
}

// DeleteEmployee performs the deletion of the employee.
func (store *SQLStore) DeleteEmployeeTx(ctx context.Context, arg DeleteEmployeeTxParams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Deletes Employee Address
		err = q.DeleteEmployeeAddress(ctx, arg.ID)
		if err != nil {
			fmt.Println("error deleting employee address", err)
			return err
		}

		// Deletes Employee Facility
		err = q.DeleteEmployeeFacilityByEmployeeId(ctx, arg.ID)
		if err != nil {
			fmt.Println("error deleting employee-facility", err)
			return err
		}

		// Deletes Employee
		err = q.DeleteEmployee(ctx, DeleteEmployeeParams(arg))

		if err != nil {
			fmt.Println("error deleting employee", err)
			return err
		}

		return err
	})

	return err
}
