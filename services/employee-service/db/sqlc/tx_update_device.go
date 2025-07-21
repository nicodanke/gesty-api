package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// UpdateDeviceTxParams contains the input parameters to update a device
type UpdateDeviceTxParams struct {
	AccountID  int64       `json:"account_id"`
	ID         int64       `json:"id"`
	Name       pgtype.Text `json:"name"`
	Enabled    pgtype.Bool `json:"enabled"`
	FacilityID pgtype.Int8 `json:"facility_id"`
	ActionIDs  []int64     `json:"action_ids"`
}

// UpdateDeviceTxResult is the result of the device update
type UpdateDeviceTxResult struct {
	Device    Device  `json:"device"`
	ActionIDs []int64 `json:"action_ids"`
}

// UpdateDeviceTx performs the update of the device.
func (store *SQLStore) UpdateDeviceTx(ctx context.Context, arg UpdateDeviceTxParams) (UpdateDeviceTxResult, error) {
	var result UpdateDeviceTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Updates Device
		result.Device, err = q.UpdateDevice(ctx, UpdateDeviceParams{
			AccountID:  arg.AccountID,
			ID:         arg.ID,
			Name:       arg.Name,
			Enabled:    arg.Enabled,
			FacilityID: arg.FacilityID,
			UpdatedAt: pgtype.Timestamptz{
				Time:  time.Now().UTC(),
				Valid: true,
			},
		})

		if err != nil {
			fmt.Println("error updating device", err)
			return err
		}

		if arg.ActionIDs != nil {
			err = q.DeleteDeviceActionByDeviceId(ctx, arg.ID)
			if err != nil {
				fmt.Println("error deleting device action", err)
				return err
			}

			// Assign all actions to device
			for _, value := range arg.ActionIDs {
				_, err = q.CreateDeviceAction(ctx, CreateDeviceActionParams{
					DeviceID: arg.ID,
					ActionID: value,
				})
				if err != nil {
					fmt.Println("error creating device action", err)
					return err
				}
			}

			result.ActionIDs = arg.ActionIDs
		} else {
			actions, err := q.GetDeviceActionsByDeviceId(ctx, arg.ID)
			if err != nil {
				fmt.Println("error getting device actions", err)
				return err
			}

			actionIDs := make([]int64, len(actions))
			for i, action := range actions {
				actionIDs[i] = action.ActionID
			}

			result.ActionIDs = actionIDs
		}

		return err
	})

	return result, err
}
