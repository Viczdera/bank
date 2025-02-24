// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: transfers.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers(
    from_account,
    to_account,
    amount
) VALUES (
    $1, $2, $3
) RETURNING _id, from_account, to_account, amount, created_at
`

type CreateTransferParams struct {
	FromAccount int64 `json:"from_account"`
	ToAccount   int64 `json:"to_account"`
	Amount      int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.queryRow(ctx, q.createTransferStmt, createTransfer, arg.FromAccount, arg.ToAccount, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccount,
		&i.ToAccount,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT _id, from_account, to_account, amount, created_at FROM transfers
WHERE _id = $1
LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, ID int64) (Transfer, error) {
	row := q.queryRow(ctx, q.getTransferStmt, getTransfer, ID)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccount,
		&i.ToAccount,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
