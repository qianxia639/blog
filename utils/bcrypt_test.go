package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {
	password := RandomString(6)

	hashPassword, err := Encrypt(password)
	require.NoError(t, err)

	err = Decrypt(password, hashPassword)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = Decrypt(wrongPassword, hashPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashPassword2, err := Encrypt(password)
	require.NoError(t, err)
	require.NotEqual(t, hashPassword, hashPassword2)
}
