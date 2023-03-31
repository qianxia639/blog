package db

import (
	"Blog/utils"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func createRandomUser(t *testing.T) User {
	username := utils.RandomString(6)
	pwd := utils.RandomString(6)

	hashPwd, err := utils.Encrypt(pwd)
	require.NoError(t, err)

	email := fmt.Sprintf("%s@email.com", username)

	arg := CreateUserParams{
		Username:     username,
		Email:        email,
		Nickname:     email,
		Password:     hashPwd,
		RegisterTime: time.Now(),
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

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	testCases := []struct {
		name  string
		logic func(*Queries) (User, error)
	}{
		{
			name: "Get User Username",
			logic: func(q *Queries) (User, error) {
				return q.GetUser(context.Background(), user1.Username)
			},
		},
		{
			name: "Get User Email",
			logic: func(q *Queries) (User, error) {
				return q.GetUser(context.Background(), user1.Email)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			user2, err := tc.logic(testQueries)
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
		name  string
		logic func(*Queries) (User, error)
	}{
		{
			name: "Username Nil",
			logic: func(q *Queries) (User, error) {
				return q.GetUser(context.Background(), "")
			},
		},
		{
			name: "Email Nil",
			logic: func(q *Queries) (User, error) {
				return q.GetUser(context.Background(), "")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			user, err := tc.logic(testQueries)
			require.Error(t, err)
			require.Empty(t, user)
		})
	}
}

func TestUpdateUserOnlyNickname(t *testing.T) {
	oldUser := createRandomUser(t)

	newNickname := utils.RandomString(6)

	newUser, err := testQueries.UpdateUser(ctx, UpdateUserParams{
		Username: oldUser.Username,
		Nickname: sql.NullString{
			String: newNickname,
			Valid:  true,
		},
	})
	require.NoError(t, err)

	require.NotEqual(t, oldUser.Nickname, newUser.Nickname)
	require.Equal(t, newNickname, newUser.Nickname)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)

	email := fmt.Sprintf("%s@email.com", utils.RandomString(6))

	newUser, err := testQueries.UpdateUser(ctx, UpdateUserParams{
		Username: oldUser.Username,
		Email: sql.NullString{
			String: email,
			Valid:  true,
		},
	})
	require.NoError(t, err)

	require.NotEqual(t, oldUser.Email, newUser.Email)
	require.Equal(t, email, newUser.Email)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)

	hashPwd, err := utils.Encrypt(utils.RandomString(6))
	require.NoError(t, err)

	newUser, err := testQueries.UpdateUser(ctx, UpdateUserParams{
		Username: oldUser.Username,
		Password: sql.NullString{
			String: hashPwd,
			Valid:  true,
		},
	})
	require.NoError(t, err)

	require.NotEqual(t, oldUser.Password, newUser.Password)
	require.Equal(t, hashPwd, newUser.Password)
}

func TestUpdateUserOnlyAvatar(t *testing.T) {
	oldUser := createRandomUser(t)

	avatar := fmt.Sprintf("%s.jpg", utils.RandomString(6))

	newUser, err := testQueries.UpdateUser(ctx, UpdateUserParams{
		Username: oldUser.Username,
		Avatar: sql.NullString{
			String: avatar,
			Valid:  true,
		},
	})
	require.NoError(t, err)

	require.NotEqual(t, oldUser.Avatar, newUser.Avatar)
	require.Equal(t, avatar, newUser.Avatar)
}

func TestUpdateUserAll(t *testing.T) {
	oldUser := createRandomUser(t)

	nickname := utils.RandomString(6)
	email := fmt.Sprintf("%s@email.com", utils.RandomString(6))

	hashPwd, err := utils.Encrypt(utils.RandomString(6))
	require.NoError(t, err)

	avatar := fmt.Sprintf("%s.jpg", utils.RandomString(6))

	newUser, err := testQueries.UpdateUser(ctx, UpdateUserParams{
		Username: oldUser.Username,
		Nickname: sql.NullString{
			String: nickname,
			Valid:  true,
		},
		Email: sql.NullString{
			String: email,
			Valid:  true,
		},
		Password: sql.NullString{
			String: hashPwd,
			Valid:  true,
		},
		Avatar: sql.NullString{
			String: avatar,
			Valid:  true,
		},
	})
	require.NoError(t, err)

	require.Equal(t, email, newUser.Email)
	require.Equal(t, hashPwd, newUser.Password)
	require.Equal(t, nickname, newUser.Nickname)
	require.Equal(t, avatar, newUser.Avatar)
}
