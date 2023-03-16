package db

import (
	"Blog/utils"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	username := utils.RandomString(6)
	pwd := utils.RandomString(6)

	hashPwd, err := utils.Encrypt(pwd)
	require.NoError(t, err)

	email := fmt.Sprintf("%s@email.com", username)

	arg := CreateUserParams{
		Username: username,
		Email:    email,
		Nickname: email,
		Password: hashPwd,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, username, user.Username)
	require.Equal(t, hashPwd, user.Password)
	require.Equal(t, email, user.Email)
	require.Equal(t, email, user.Nickname)
	require.False(t, user.RegisterTime.IsZero())

	return user
}

func TestCreeateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	testCases := []struct {
		name string
		fn   func(*Queries) (User, error)
	}{
		{
			name: "Get User Username",
			fn: func(q *Queries) (User, error) {
				return q.GetUser(context.Background(), user1.Username)
			},
		},
		{
			name: "Get User Email",
			fn: func(q *Queries) (User, error) {
				return q.GetUser(context.Background(), user1.Email)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			user2, err := tc.fn(testQueries)
			require.NoError(t, err)

			require.Equal(t, user1.ID, user2.ID)
			require.Equal(t, user1.Username, user2.Username)
			require.Equal(t, user1.Password, user2.Password)
			require.Equal(t, user1.Email, user2.Email)
			require.Equal(t, user1.Nickname, user2.Nickname)
			require.WithinDuration(t, user1.RegisterTime, user2.RegisterTime, time.Second)
		})
	}
}

func TestGetUserEmpty(t *testing.T) {
	testCases := []struct {
		name string
		fn   func(*Queries) (User, error)
	}{
		{
			name: "Username Nil",
			fn: func(q *Queries) (User, error) {
				return q.GetUser(context.Background(), "")
			},
		},
		{
			name: "Email Nil",
			fn: func(q *Queries) (User, error) {
				return q.GetUser(context.Background(), "")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			user, err := tc.fn(testQueries)
			require.Error(t, err)
			require.Empty(t, user)
		})
	}
}
