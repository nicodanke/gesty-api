package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

// CreateEmployeeTxParams contains the input parameters to create a employee
type CreateEmployeeTxParams struct {
	AccountID         int64   `json:"account_id"`
	Name              string  `json:"name"`
	Lastname          string  `json:"lastname"`
	Email             string  `json:"email"`
	Phone             string  `json:"phone"`
	Gender            string  `json:"gender"`
	RealID            string  `json:"real_id"`
	FiscalID          string  `json:"fiscal_id"`
	AddressCountry    string  `json:"address_country"`
	AddressState      string  `json:"address_state"`
	AddressSubState   string  `json:"address_sub_state"`
	AddressStreet     string  `json:"address_street"`
	AddressNumber     string  `json:"address_number"`
	AddressUnit       string  `json:"address_unit"`
	AddressPostalcode string  `json:"address_postalcode"`
	AddressLat        float64 `json:"address_lat"`
	AddressLng        float64 `json:"address_lng"`
	FacilityIDs       []int64 `json:"facility_ids"`
}

// CreateEmployeeTxResult is the result of the employee creation
type CreateEmployeeTxResult struct {
	Employee        Employee        `json:"employee"`
	EmployeeAddress EmployeeAddress `json:"employee_address"`
	FacilityIDs     []int64         `json:"facility_ids"`
}

// CreateEmployeeTx performs the creation of the employee and the facilities.
func (store *SQLStore) CreateEmployeeTx(ctx context.Context, arg CreateEmployeeTxParams) (CreateEmployeeTxResult, error) {
	var result CreateEmployeeTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Creates Employee
		result.Employee, err = q.CreateEmployee(ctx, CreateEmployeeParams{
			AccountID: arg.AccountID,
			Name:      arg.Name,
			Lastname:  arg.Lastname,
			Email:     arg.Email,
			Phone:     arg.Phone,
			Gender:    arg.Gender,
			RealID:    arg.RealID,
			FiscalID:  arg.FiscalID,
		})

		if err != nil {
			fmt.Println("error creating employee", err)
			return err
		}

		// Creates Employee Address
		result.EmployeeAddress, err = q.CreateEmployeeAddress(ctx, CreateEmployeeAddressParams{
			EmployeeID: result.Employee.ID,
			Country:    arg.AddressCountry,
			State:      arg.AddressState,
			SubState: pgtype.Text{
				String: arg.AddressSubState,
				Valid:  arg.AddressSubState != "",
			},
			Street: arg.AddressStreet,
			Number: arg.AddressNumber,
			Unit: pgtype.Text{
				String: arg.AddressUnit,
				Valid:  arg.AddressUnit != "",
			},
			PostalCode: arg.AddressPostalcode,
			Lat: pgtype.Float8{
				Float64: arg.AddressLat,
				Valid:   true,
			},
			Lng: pgtype.Float8{
				Float64: arg.AddressLng,
				Valid:   true,
			},
		})

		// Assign all facilities to employee
		for _, value := range arg.FacilityIDs {
			_, err = q.CreateEmployeeFacility(ctx, CreateEmployeeFacilityParams{
				EmployeeID: result.Employee.ID,
				FacilityID: int64(value),
			})
			if err != nil {
				return err
			}
		}

		result.FacilityIDs = arg.FacilityIDs

		return err
	})

	return result, err
}
