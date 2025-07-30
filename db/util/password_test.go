package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(6)
	hashedPass, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPass)

	checkPasswordErr := CheckPassword(password, hashedPass)
	require.NoError(t, checkPasswordErr)

	password2 := RandomString(6)
	checkPasswordErr2 := CheckPassword(password2, hashedPass)

	require.EqualError(t, checkPasswordErr2, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPass2, _ := HashPassword(password2)
	require.NotEqual(t, hashedPass, hashedPass2)

}
