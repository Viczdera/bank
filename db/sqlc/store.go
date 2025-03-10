package db

import (
	"context"
	"database/sql"
	"fmt"
)

// store provides all functions to execute a transaction in a db
// queries Struct is extended. This is know as composition. Kinda like inheritance
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}

	//New function was created by sqlc. Creates and returns a new query object
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	//begin a new transaction
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	//create a new Queries object tied to the transaction
	newQ := New(tx)

	//execute provided function within the transaction
	err = fn(newQ)

	if err != nil {
		//return err if transaction rolback fails
		if errRb := tx.Rollback(); errRb != nil {
			// return errRb
			return fmt.Errorf("tx err: %v, rb err: %v", err, errRb)
		}
		return err
	}

	//commit transaction if all is successful
	return tx.Commit()

}

//Transfer function to transfer money from one account to another
//It will create a transfer record, add account entries and update the accounts

type TransferTxParams struct {
	FromAccount int64 `json:"from_account"`
	ToAccount   int64 `json:"to_account"`
	Amount      int64 `json:"amount"`
}

// structure of result after the transaction is complete
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
}

// the result returns the result and an error
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		//we're accessing the arg and result from the outer function. This makes it a closure. Used when we want to get the result of a callback function cause the callback function itself does not know the type of the result

		var err error

		//make hte transfer
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccount: arg.FromAccount,
			ToAccount:   arg.ToAccount,
			Amount:      arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccount,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccount,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccount,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccount,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
