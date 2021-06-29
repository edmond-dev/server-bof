package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		UserID:    "kkjdo934jkjd",
		FirstName: "a",
		LastName:  "aa",
		Email:     "a@a.com",
		Password:  "aaa",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserID, user)

	require.Equal(t, arg.FirstName, arg.FirstName)
	require.Equal(t, arg.LastName, arg.LastName)
	require.Equal(t, arg.Email, arg.Email)

	//require.NotZero(t, )
}
