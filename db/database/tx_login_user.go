package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/google/uuid"
)

type LoginUserTxParams struct {
	Username        string
	Password        string
	DeviceName      string
	DeviceHash      string
	OperatingSystem string
	MacAddress      string
	ClientIp        string
	UserAgent       string
}

func (gojo *SQLGojo) LoginUserTx(ctx context.Context, arg LoginUserTxParams) (User, error) {
	var user User
	var err error

	err = gojo.execTx(ctx, func(q *Queries) error {
		user, err = q.GetUserByUsername(ctx, arg.Username)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		err = utils.CheckPassword(arg.Password, user.HashedPassword)
		if err != nil {
			return ErrInccorectPassword
		}

		device, err := q.CreateDevice(ctx, CreateDeviceParams{
			ID:              uuid.New(),
			DeviceName:      arg.DeviceName,
			DeviceHash:      arg.DeviceHash,
			OperatingSystem: arg.OperatingSystem,
			MacAddress:      arg.MacAddress,
			ClientIp:        arg.ClientIp,
			UserAgent:       arg.UserAgent,
		})
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if device.IsBanned {
			return ErrFailedPrecondition
		}

		if err = q.CreateUserDevice(ctx, CreateUserDeviceParams{
			UserID:   user.ID,
			DeviceID: device.ID,
		}); err != nil {
			ErrorSQL(err)
			return err
		}

		return err
	})

	return user, err
}
