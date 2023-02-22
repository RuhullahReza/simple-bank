package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/RuhullahReza/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account{
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner: user.Username,
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t,err)
	require.NotEmpty(t,account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	foundAccount, err := testQueries.GetAccount(context.Background(),newAccount.ID)
	require.NoError(t,err)
	require.NotEmpty(t,foundAccount)

	require.Equal(t, newAccount.ID, foundAccount.ID)
	require.Equal(t, newAccount.Owner, foundAccount.Owner)
	require.Equal(t, newAccount.Balance, foundAccount.Balance)
	require.Equal(t, newAccount.Currency, foundAccount.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, foundAccount.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID: newAccount.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,updatedAccount)

	require.Equal(t, newAccount.ID, updatedAccount.ID)
	require.Equal(t, newAccount.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, newAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)

	deletedAccount, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.Error(t,err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountParams{
		Owner: lastAccount.Owner,
		Limit: 5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _,account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}