package db

import (
	"context"
	"fmt"
)

// DeleteDeviceTxParams contains the input parameters to delete a device
type DeleteDeviceTxParams struct {
	AccountID int64 `json:"account_id"`
	ID        int64 `json:"id"`
}

// DeleteDeviceTx performs the deletion of the device.
func (store *SQLStore) DeleteDeviceTx(ctx context.Context, arg DeleteDeviceTxParams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Deletes Device Action
		err = q.DeleteDeviceActionByDeviceId(ctx, arg.ID)
		if err != nil {
			fmt.Println("error deleting device action", err)
			return err
		}

		// Deletes Device Health
		err = q.DeleteDeviceHealth(ctx, arg.ID)
		if err != nil {
			fmt.Println("error deleting device health", err)
			return err
		}

		// Deletes Device
		err = q.DeleteDevice(ctx, DeleteDeviceParams(arg))

		if err != nil {
			fmt.Println("error deleting device", err)
			return err
		}

		return err
	})

	return err
}
