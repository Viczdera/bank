package db

import (
	"context"
	"testing"
	"time"

	"github.com/Viczdera/bank/db/util"
	"github.com/stretchr/testify/require"
)

func createTestUser(t *testing.T) User {
	args := CreateUserParams{
		Username:       util.RandomOwner(),
		PasswordHashed: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.PasswordHashed, user.PasswordHashed)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user

}

func TestCreateUser(t *testing.T) {
	createTestUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createTestUser(t)
	userFind, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, userFind)

	require.Equal(t, user1.Username, userFind.Username)
	require.Equal(t, user1.PasswordHashed, userFind.PasswordHashed)
	require.Equal(t, user1.FullName, userFind.FullName)
	require.Equal(t, user1.Email, userFind.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, userFind.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, userFind.CreatedAt, time.Second)
}
