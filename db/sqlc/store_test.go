package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createTestAccount(t)
	account2 := createTestAccount(t)
	fmt.Println(">>before:", account1.Balance, account2.Balance)

	//do it 6 times
	n := 6
	randAmount := int64(10)
	// make() is a built-in Go function that allocates and initializes memory
	// here it creates a map that stores bool values with int keys
	existedTransfers := make(map[int]bool)
	// make() creates buffered channels with capacity n
	// channels are used to communicate between goroutines
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

		//check entries
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.EqualValues(t, account2.ID, toEntry.AccountID)
		require.EqualValues(t, randAmount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//TODO:check account balances
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.EqualValues(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.EqualValues(t, account2.ID, toAccount.ID)

		fmt.Println(">>tx:", fromAccount.Balance, toAccount.Balance)

		//differences
		fromAccountDiff := account1.Balance - fromAccount.Balance
		toAccountDiff := toAccount.Balance - account2.Balance
		require.EqualValues(t, fromAccountDiff, toAccountDiff)
		require.True(t, fromAccountDiff > 0)
		require.True(t, fromAccountDiff%randAmount == 0)

		//check amount of times transactions were made
		tCount := int(fromAccountDiff / randAmount)
		require.True(t, tCount >= 1 && tCount <= n)
		require.NotContains(t, existedTransfers, tCount)
		existedTransfers[tCount] = true

	}

	//check the final balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">>after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.EqualValues(t, account1.Balance-(int64(n)*randAmount), updatedAccount1.Balance)
	require.EqualValues(t, account2.Balance+(int64(n)*randAmount), updatedAccount2.Balance)
}
