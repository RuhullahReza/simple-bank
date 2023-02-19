package db

import (
	"context"
	"testing"
	"time"

	"github.com/RuhullahReza/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T){
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transferData := createRandomTransfer(t, account1, account2)

	transfer, err := testQueries.GetTransfer(context.Background(), transferData.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transferData.ID, transfer.ID)
	require.Equal(t, transferData.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transferData.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transferData.Amount, transfer.Amount)
	require.WithinDuration(t, transferData.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T){
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t,account1,account2)
		createRandomTransfer(t,account2,account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID: account1.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}

}