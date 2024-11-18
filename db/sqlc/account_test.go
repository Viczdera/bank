package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Viczdera/bank/db/util"
	"github.com/stretchr/testify/require"
)

func createTestAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	//instead of if-else, use testify: https://github.com/stretchr/testify

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account

}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	acct1 := createTestAccount(t)
	acctFind, err := testQueries.GetAccount(context.Background(), acct1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acctFind)

	require.Equal(t, acct1.Balance, acctFind.Balance)
	require.Equal(t, acct1.Currency, acctFind.Currency)
	require.Equal(t, acct1.Owner, acctFind.Owner)
	require.WithinDuration(t, acct1.CreatedAt, acctFind.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acct1 := createTestAccount(t)
	args := UpdateAccountParams{
		ID:      acct1.ID,
		Balance: util.RandomBalance(),
	}

	acctUpdate, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, acctUpdate)

	require.Equal(t, args.Balance, acctUpdate.Balance)
	require.Equal(t, acct1.ID, acctUpdate.ID)
	require.Equal(t, acct1.Currency, acctUpdate.Currency)
	require.Equal(t, acct1.Owner, acctUpdate.Owner)
	require.WithinDuration(t, acct1.CreatedAt, acctUpdate.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	acct1 := createTestAccount(t)
	err := testQueries.DeleteAccount(context.Background(), acct1.ID)

	require.NoError(t, err)

	acctFind, err := testQueries.GetAccount(context.Background(), acct1.ID)

	require.Error(t, err)
	//err must match sql no rows found err
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acctFind)
}

func TestListAccount(t *testing.T) {

	for i := 0; i < 10; i++ {
		createTestAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for i := 0; i < len(accounts); i++ {
		require.NotEmpty(t, accounts[i])
	}
	//duplicate yunno
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
