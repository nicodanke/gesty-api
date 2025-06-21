package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (CreateAccountTxResult, error)
	CreateRoleTx(ctx context.Context, arg CreateRoleTxParams) (CreateRoleTxResult, error)
	UpdateRoleTx(ctx context.Context, arg UpdateRoleTxParams) (UpdateRoleTxResult, error)
	DeleteRoleTx(ctx context.Context, arg DeleteRoleTxParams) error
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