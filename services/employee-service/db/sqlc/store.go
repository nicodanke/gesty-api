package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) error
	CreateFacilityTx(ctx context.Context, arg CreateFacilityTxParams) (CreateFacilityTxResult, error)
	UpdateFacilityTx(ctx context.Context, arg UpdateFacilityTxParams) (UpdateFacilityTxResult, error)
	DeleteFacilityTx(ctx context.Context, arg DeleteFacilityTxParams) error
	CreateEmployeeTx(ctx context.Context, arg CreateEmployeeTxParams) (CreateEmployeeTxResult, error)
	UpdateEmployeeTx(ctx context.Context, arg UpdateEmployeeTxParams) (UpdateEmployeeTxResult, error)
	DeleteEmployeeTx(ctx context.Context, arg DeleteEmployeeTxParams) error
	CreateDeviceTx(ctx context.Context, arg CreateDeviceTxParams) (CreateDeviceTxResult, error)
	UpdateDeviceTx(ctx context.Context, arg UpdateDeviceTxParams) (UpdateDeviceTxResult, error)
	DeleteDeviceTx(ctx context.Context, arg DeleteDeviceTxParams) error
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
