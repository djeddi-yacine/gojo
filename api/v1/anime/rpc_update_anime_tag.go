package av1

import (
	"context"
	"errors"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeServer) UpdateAnimeTag(ctx context.Context, req *apbv1.UpdateAnimeTagRequest) (*apbv1.UpdateAnimeTagResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update anime tag")
	}

	if violations := validateUpdateAnimeTagRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.UpdateAnimeTagParams{
		ID: req.GetTagID(),
		Tag: pgtype.Text{
			String: req.GetTag(),
			Valid:  req.Tag != nil,
		},
	}

	data, err := server.gojo.UpdateAnimeTag(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to update anime tag", err)
	}

	res := &apbv1.UpdateAnimeTagResponse{
		AnimeTag: &apbv1.AnimeTag{
			ID:        data.ID,
			Tag:       data.Tag,
			CreatedAt: timestamppb.New(data.CreatedAt),
		},
	}

	return res, nil
}

func validateUpdateAnimeTagRequest(req *apbv1.UpdateAnimeTagRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetTagID()); err != nil {
		violations = append(violations, shv1.FieldViolation("tagID", err))
	}

	if req.Tag != nil {
		if err := utils.ValidateString(req.GetTag(), 2, 500); err != nil {
			violations = append(violations, shv1.FieldViolation("tag", err))
		}
	} else {
		violations = append(violations, shv1.FieldViolation("tag", errors.New("put the new tag")))
	}

	return violations
}
