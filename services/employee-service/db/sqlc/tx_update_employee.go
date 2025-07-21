package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// UpdateEmployeeTxParams contains the input parameters to update an employee
type UpdateEmployeeTxParams struct {
	AccountID       int64         `json:"account_id"`
	ID              int64         `json:"id"`
	Name            pgtype.Text   `json:"name"`
	Lastname        pgtype.Text   `json:"lastname"`
	Email           pgtype.Text   `json:"email"`
	Phone           pgtype.Text   `json:"phone"`
	Gender          pgtype.Text   `json:"gender"`
	RealId          pgtype.Text   `json:"real_id"`
	FiscalId        pgtype.Text   `json:"fiscal_id"`
	AddressCountry  pgtype.Text   `json:"address_country"`
	AddressState    pgtype.Text   `json:"address_state"`
	AddressSubState pgtype.Text   `json:"address_sub_state"`
	AddressStreet   pgtype.Text   `json:"address_street"`
	AddressNumber   pgtype.Text   `json:"address_number"`
	AddressUnit     pgtype.Text   `json:"address_unit"`
	AddressZipCode  pgtype.Text   `json:"address_ZipCode"`
	AddressLat      pgtype.Float8 `json:"address_lat"`
	AddressLng      pgtype.Float8 `json:"address_lng"`
	FacilityIds     []int64       `json:"facility_ids"`
}

// UpdateEmployeeTxResult is the result of the employee update
type UpdateEmployeeTxResult struct {
	Employee        Employee        `json:"employee"`
	EmployeeAddress EmployeeAddress `json:"employee_address"`
	FacilityIds     []int64         `json:"facility_ids"`
}

// UpdateEmployeeTx performs the update of the employee.
func (store *SQLStore) UpdateEmployeeTx(ctx context.Context, arg UpdateEmployeeTxParams) (UpdateEmployeeTxResult, error) {
	var result UpdateEmployeeTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Updates Employee
		result.Employee, err = q.UpdateEmployee(ctx, UpdateEmployeeParams{
			AccountID: arg.AccountID,
			ID:        arg.ID,
			Name:      arg.Name,
			Lastname:  arg.Lastname,
			Email:     arg.Email,
			Phone:     arg.Phone,
			Gender:    arg.Gender,
			RealID:    arg.RealId,
			FiscalID:  arg.FiscalId,
			UpdatedAt: pgtype.Timestamptz{
				Time:  time.Now().UTC(),
				Valid: true,
			},
		})

		if err != nil {
			fmt.Println("error updating employee", err)
			return err
		}

		if arg.AddressCountry.Valid || arg.AddressState.Valid || arg.AddressSubState.Valid || arg.AddressStreet.Valid || arg.AddressNumber.Valid || arg.AddressUnit.Valid || arg.AddressZipCode.Valid || arg.AddressLat.Valid || arg.AddressLng.Valid {
			result.EmployeeAddress, err = q.UpdateEmployeeAddress(ctx, UpdateEmployeeAddressParams{
				EmployeeID: arg.ID,
				Country:    arg.AddressCountry,
				State:      arg.AddressState,
				SubState:   arg.AddressSubState,
				Street:     arg.AddressStreet,
				Number:     arg.AddressNumber,
				Unit:       arg.AddressUnit,
				ZipCode:    arg.AddressZipCode,
				Lat:        arg.AddressLat,
				Lng:        arg.AddressLng,
			})
		} else {
			result.EmployeeAddress, err = q.GetEmployeeAddressByEmployeeId(ctx, arg.ID)
		}

		if arg.FacilityIds != nil {
			err = q.DeleteEmployeeFacilityByEmployeeId(ctx, arg.ID)
			if err != nil {
				fmt.Println("error deleting employee-facility", err)
				return err
			}

			// Assign all permissions to role
			for _, value := range arg.FacilityIds {
				_, err = q.CreateEmployeeFacility(ctx, CreateEmployeeFacilityParams{
					EmployeeID: arg.ID,
					FacilityID: value,
				})
				if err != nil {
					fmt.Println("error creating role permissions", err)
					return err
				}
			}
			result.FacilityIds = arg.FacilityIds
		} else {
			facilities, err := q.GetEmployeeFacilitiesByEmployeeId(ctx, arg.ID)
			if err != nil {
				fmt.Println("error getting employee facilities", err)
				return err
			}

			facilityIDs := make([]int64, len(facilities))
			for i, facility := range facilities {
				facilityIDs[i] = facility.FacilityID
			}

			result.FacilityIds = facilityIDs
		}

		return err
	})

	return result, err
}
