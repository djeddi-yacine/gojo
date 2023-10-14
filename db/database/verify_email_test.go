package db

import (
	"context"
	"testing"
	"time"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/stretchr/testify/require"
)

func createRandomVerifyEmail(t *testing.T) VerifyEmail {
	u := createRandomUser(t)
	require.NotEmpty(t, u)

	arg := CreateVerifyEmailParams{
		Username:   u.Username,
		Email:      u.Email,
		SecretCode: utils.RandomString(10),
	}

	verifyEmail, err := testGojo.CreateVerifyEmail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, verifyEmail)

	require.Equal(t, arg.Username, verifyEmail.Username)
	require.Equal(t, arg.Email, verifyEmail.Email)
	require.Equal(t, arg.SecretCode, verifyEmail.SecretCode)

	require.NotZero(t, verifyEmail.ID)
	require.NotNil(t, verifyEmail.CreatedAt)
	require.NotNil(t, verifyEmail.ExpiredAt)
	return verifyEmail
}

func TestCreateVerifyEmail(t *testing.T) {
	createRandomVerifyEmail(t)
}

func TestUpdateVerifyEmail(t *testing.T) {
	v := createRandomVerifyEmail(t)
	require.NotEmpty(t, v)

	arg := UpdateVerifyEmailParams{
		ID:         v.ID,
		SecretCode: v.SecretCode,
	}

	verifyEmail, err := testGojo.UpdateVerifyEmail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, verifyEmail)

	require.Equal(t, v.ID, verifyEmail.ID)
	require.Equal(t, v.Email, verifyEmail.Email)
	require.Equal(t, v.SecretCode, verifyEmail.SecretCode)
	require.True(t, verifyEmail.IsUsed)

	require.WithinDuration(t, v.ExpiredAt, verifyEmail.ExpiredAt, time.Second)
	require.WithinDuration(t, v.CreatedAt, verifyEmail.CreatedAt, time.Second)
}
