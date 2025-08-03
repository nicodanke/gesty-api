package db

import (
	"context"
	"errors"
	"fmt"
)

// CreateEmployeePhotoTxParams contains the input parameters to create a employee photo
type CreateEmployeePhotoTxParams struct {
	EmployeeID  int64  `json:"employee_id"`
	ImageBase64 string `json:"image_base_64"`
	VectorImage string `json:"vector_image"`
	AccountID   int64  `json:"account_id"`
}

// CreateEmployeePhotoTxResult is the result of the employee photo creation
type CreateEmployeePhotoTxResult struct {
	EmployeePhoto EmployeePhoto `json:"employee_photo"`
}

// CreateEmployeePhotoTx performs the creation of the employee photo.
func (store *SQLStore) CreateEmployeePhotoTx(ctx context.Context, arg CreateEmployeePhotoTxParams) (CreateEmployeePhotoTxResult, error) {
	var result CreateEmployeePhotoTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Get Photos to check quantity
		photos, err := q.GetEmployeePhotos(ctx, GetEmployeePhotosParams{
			EmployeeID: arg.EmployeeID,
			AccountID:  arg.AccountID,
			Limit:      100,
			Offset:     0,
		})

		if err != nil {
			return err
		}

		if len(photos) >= 10 {
			return errors.New("Employee has already 10 photos")
		}

		// Creates Employee Photo
		result.EmployeePhoto, err = q.CreateEmployeePhoto(ctx, CreateEmployeePhotoParams{
			EmployeeID:  arg.EmployeeID,
			ImageBase64: arg.ImageBase64,
			VectorImage: arg.VectorImage,
			AccountID:   arg.AccountID,
			IsProfile:   len(photos) == 0,
		})

		if err != nil {
			fmt.Println("error creating employee photo", err)
			return err
		}

		return err
	})

	return result, err
}
