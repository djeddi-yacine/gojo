package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomSession(t *testing.T) Session {
	u := createRandomUser(t)
	require.NotEmpty(t, u)

	arg := CreateSessionParams{
		ID:        uuid.New(),
		Username:  u.Username,
		IsBlocked: false,
		ExpiresAt: time.Now(),
	}

	session, err := testGojo.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.Equal(t, arg.ID, session.ID)
	require.Equal(t, arg.Username, session.Username)
	require.Equal(t, arg.IsBlocked, session.IsBlocked)
	require.WithinDuration(t, arg.ExpiresAt, session.ExpiresAt, time.Second)

	return session
}

func TestCreateSession(t *testing.T) {
	createRandomSession(t)
}

func TestGetSession(t *testing.T) {
	s := createRandomSession(t)
	require.NotEmpty(t, s)

	session, err := testGojo.GetSession(context.Background(), s.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.Equal(t, s.ID, session.ID)
	require.Equal(t, s.Username, session.Username)
	require.Equal(t, s.IsBlocked, session.IsBlocked)
	require.Equal(t, s.ClientIp, session.ClientIp)
	require.Equal(t, s.UserAgent, session.UserAgent)
	require.Equal(t, s.RefreshToken, session.RefreshToken)

	require.WithinDuration(t, s.ExpiresAt, session.ExpiresAt, time.Second)
	require.WithinDuration(t, s.CreatedAt, session.CreatedAt, time.Second)

}
