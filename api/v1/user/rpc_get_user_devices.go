package usapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	uspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/uspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *UserServer) GetUserDevices(ctx context.Context, req *uspbv1.GetUserDevicesRequest) (*uspbv1.GetUserDevicesResponse, error) {
	if violations := validateGetUserDevicesRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	user, err := server.gojo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, shv1.ApiError("failed to get user", err)
	}

	devicesIDs, err := server.gojo.ListUserDevices(ctx, db.ListUserDevicesParams{
		UserID: user.ID,
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageNumber() - 1) * req.GetPageSize(),
	})
	if err != nil {
		return nil, shv1.ApiError("failed to list devices IDs", err)
	}

	devices := make([]*uspbv1.Device, len(devicesIDs))

	for i, d := range devicesIDs {
		device, err := server.gojo.GetDevice(ctx, d)
		if err != nil {
			continue
		}

		devices[i] = &uspbv1.Device{
			ID:              device.ID.String(),
			OperatingSystem: device.OperatingSystem,
			MacAddress:      device.MacAddress,
			ClientIP:        device.ClientIp,
			UserAgent:       device.UserAgent,
			IsBanned:        device.IsBanned,
			CreatedAt:       timestamppb.New(device.CreatedAt),
		}
	}

	res := &uspbv1.GetUserDevicesResponse{
		Devices: devices,
	}

	return res, nil
}

func validateGetUserDevicesRequest(req *uspbv1.GetUserDevicesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetUserID()); err != nil {
		violations = append(violations, shv1.FieldViolation("userID", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	return violations
}
