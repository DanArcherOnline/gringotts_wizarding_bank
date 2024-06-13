package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/danarcheronline/gringotts_wizarding_bank/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomFloat(0, 1000),
		Currency: util.RandomCurrency(),
	}

	acc, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, args.Balance, acc.Balance)
	require.Equal(t, args.Currency, acc.Currency)
	require.Equal(t, args.Owner, acc.Owner)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Balance, acc2.Balance)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, acc1.CreatedAt, acc2.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	updateAccountParams := UpdateAccountParams{
		ID:      acc1.ID,
		Balance: util.RandomFloat(1001, 2000),
	}

	acc2, err := testQueries.UpdateAccount(context.Background(), updateAccountParams)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc1.ID, acc2.ID)
	require.NotEqual(t, acc1.Balance, acc2.Balance)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, acc1.CreatedAt, acc2.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	err1 := testQueries.DeleteAccount(context.Background(), acc1.ID)
	acc2, err2 := testQueries.GetAccount(context.Background(), acc1.ID)

	require.NoError(t, err1)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, acc2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	listAccountsParams := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), listAccountsParams)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
