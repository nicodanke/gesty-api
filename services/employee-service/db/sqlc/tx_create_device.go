package db

import (
	"context"
	"fmt"
)

// CreateDeviceTxParams contains the input parameters to create a device
type CreateDeviceTxParams struct {
	AccountID  int64   `json:"account_id"`
	Name       string  `json:"name"`
	Enabled    bool    `json:"enabled"`
	Password   string  `json:"password"`
	FacilityID int64   `json:"facility_id"`
	ActionIDs  []int64 `json:"action_ids"`
}

// CreateDeviceTxResult is the result of the device creation
type CreateDeviceTxResult struct {
	Device    Device  `json:"device"`
	ActionIDs []int64 `json:"action_ids"`
}

// CreateDeviceTx performs the creation of the device and the actions.
func (store *SQLStore) CreateDeviceTx(ctx context.Context, arg CreateDeviceTxParams) (CreateDeviceTxResult, error) {
	var result CreateDeviceTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Creates Device
		result.Device, err = q.CreateDevice(ctx, CreateDeviceParams{
			AccountID:  arg.AccountID,
			Name:       arg.Name,
			Enabled:    arg.Enabled,
			Password:   arg.Password,
			FacilityID: arg.FacilityID,
		})

		if err != nil {
			fmt.Println("error creating device", err)
			return err
		}

		// Assign all actions to device
		for _, value := range arg.ActionIDs {
			_, err = q.CreateDeviceAction(ctx, CreateDeviceActionParams{
				DeviceID: result.Device.ID,
				ActionID: int64(value),
			})
			if err != nil {
				return err
			}
		}

		result.ActionIDs = arg.ActionIDs

		return err
	})

	return result, err
}
