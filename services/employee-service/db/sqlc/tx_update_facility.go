package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// UpdateFacilityTxParams contains the input parameters to update a facility
type UpdateFacilityTxParams struct {
	AccountID       int64         `json:"account_id"`
	ID              int64         `json:"id"`
	Name            pgtype.Text   `json:"name"`
	Description     pgtype.Text   `json:"description"`
	OpenTime        pgtype.Time   `json:"open_time"`
	CloseTime       pgtype.Time   `json:"close_time"`
	AddressCountry  pgtype.Text   `json:"address_country"`
	AddressState    pgtype.Text   `json:"address_state"`
	AddressSubState pgtype.Text   `json:"address_sub_state"`
	AddressStreet   pgtype.Text   `json:"address_street"`
	AddressNumber   pgtype.Text   `json:"address_number"`
	AddressUnit     pgtype.Text   `json:"address_unit"`
	AddressZipCode  pgtype.Text   `json:"address_ZipCode"`
	AddressLat      pgtype.Float8 `json:"address_lat"`
	AddressLng      pgtype.Float8 `json:"address_lng"`
}

// UpdateFacilityTxResult is the result of the facility update
type UpdateFacilityTxResult struct {
	Facility        Facility        `json:"facility"`
	FacilityAddress FacilityAddress `json:"facility_address"`
}

// UpdateFacilityTx performs the update of the facility.
func (store *SQLStore) UpdateFacilityTx(ctx context.Context, arg UpdateFacilityTxParams) (UpdateFacilityTxResult, error) {
	var result UpdateFacilityTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Updates Facility
		result.Facility, err = q.UpdateFacility(ctx, UpdateFacilityParams{
			AccountID:   arg.AccountID,
			ID:          arg.ID,
			Name:        arg.Name,
			Description: arg.Description,
			OpenTime:    arg.OpenTime,
			CloseTime:   arg.CloseTime,
			UpdatedAt: pgtype.Timestamptz{
				Time:  time.Now().UTC(),
				Valid: true,
			},
		})

		if err != nil {
			fmt.Println("error updating role", err)
			return err
		}

		if arg.AddressCountry.Valid || arg.AddressState.Valid || arg.AddressSubState.Valid || arg.AddressStreet.Valid || arg.AddressNumber.Valid || arg.AddressUnit.Valid || arg.AddressZipCode.Valid || arg.AddressLat.Valid || arg.AddressLng.Valid {
			result.FacilityAddress, err = q.UpdateFacilityAddress(ctx, UpdateFacilityAddressParams{
				FacilityID: arg.ID,
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
			result.FacilityAddress, err = q.GetFacilityAddressByFacilityID(ctx, arg.ID)
		}

		return err
	})

	return result, err
}
