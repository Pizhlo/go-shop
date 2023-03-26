package db

import (
	"context"
	"github.com/Pizhlo/go-shop/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: util.RandomString(10),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, user)

	require.NotZero(t, user.ID)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, []int32(nil), user.Favourites)
}
