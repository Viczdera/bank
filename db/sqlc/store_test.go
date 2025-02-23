package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createTestAccount(t)
	account2 := createTestAccount(t)

	//do it for 6 account
	n := 6
	randAmount := int64(10)

	//channels for goroutines
	results := make(chan (TransferTxResult), n)
	errors := make(chan (error), n)

	for i := 0; i < n; i++ {

		go func(i int) {

			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccount: account1.ID,
				ToAccount:   account2.ID,
				Amount:      randAmount,
			})

			results <- result
			errors <- err
		}(i)

	}

	for i := 0; i < n; i++ {
		err := <-errors
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//if result is not empty then there's a transfer record. check records
		transfer := result.Transfer
		require.EqualValues(t, account1.ID, transfer.FromAccount)
		require.EqualValues(t, account2.ID, transfer.ToAccount)
		require.EqualValues(t, randAmount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		//check error in transfer
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//finally, check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.EqualValues(t, account1.ID, fromEntry.AccountID)
		require.EqualValues(t, -randAmount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		//finally, check entries
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.EqualValues(t, account2.ID, toEntry.AccountID)
		require.EqualValues(t, randAmount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//TODO:check account balances
	}

}
