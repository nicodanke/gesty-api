package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateFacilityTx(ctx context.Context, arg CreateFacilityTxParams) (CreateFacilityTxResult, error)
	UpdateFacilityTx(ctx context.Context, arg UpdateFacilityTxParams) (UpdateFacilityTxResult, error)
	DeleteFacilityTx(ctx context.Context, arg DeleteFacilityTxParams) error
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
