package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/durationpb"
)

// CreateFacilityTxParams contains the input parameters to create a facility
type CreateFacilityTxParams struct {
	AccountID         int64                `json:"account_id"`
	Name              string               `json:"name"`
	Description       string               `json:"description"`
	OpenTime          *durationpb.Duration `json:"open_time"`
	CloseTime         *durationpb.Duration `json:"close_time"`
	AddressCountry    string               `json:"address_country"`
	AddressState      string               `json:"address_state"`
	AddressSubState   string               `json:"address_sub_state"`
	AddressStreet     string               `json:"address_street"`
	AddressNumber     string               `json:"address_number"`
	AddressUnit       string               `json:"address_unit"`
	AddressPostalcode string               `json:"address_postalcode"`
	AddressLat        float64              `json:"address_lat"`
	AddressLng        float64              `json:"address_lng"`
}

// CreateFacilityTxResult is the result of the facility creation
type CreateFacilityTxResult struct {
	Facility        Facility        `json:"facility"`
	FacilityAddress FacilityAddress `json:"facility_address"`
}

// CreateFacilityTx performs the creation of the facility.
func (store *SQLStore) CreateFacilityTx(ctx context.Context, arg CreateFacilityTxParams) (CreateFacilityTxResult, error) {
	var result CreateFacilityTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Creates Facility
		result.Facility, err = q.CreateFacility(ctx, CreateFacilityParams{
			AccountID: arg.AccountID,
			Name:      arg.Name,
			Description: pgtype.Text{
				String: arg.Description,
				Valid:  arg.Description != "",
			},
			OpenTime: pgtype.Time{
				Microseconds: arg.OpenTime.AsDuration().Microseconds(),
				Valid:        arg.OpenTime != nil,
			},
			CloseTime: pgtype.Time{
				Microseconds: arg.CloseTime.AsDuration().Microseconds(),
				Valid:        arg.CloseTime != nil,
			},
		})

		if err != nil {
			fmt.Println("error creating facility", err)
			return err
		}

		// Create Facility Address
		result.FacilityAddress, err = q.CreateFacilityAddress(ctx, CreateFacilityAddressParams{
			FacilityID: result.Facility.ID,
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

		if err != nil {
			fmt.Println("error creating facility address", err)
			return err
		}

		return err
	})

	return result, err
}
