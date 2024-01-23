package av1

import (
	"context"

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

func (server *AnimeServer) UpdateAnimeImage(ctx context.Context, req *apbv1.UpdateAnimeImageRequest) (*apbv1.UpdateAnimeImageResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update anime character")
	}

	if violations := validateUpdateAnimeImageRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.UpdateAnimeImageParams{
		ID: req.GetImageID(),
		ImageHost: pgtype.Text{
			String: req.GetHost(),
			Valid:  req.Host != nil,
		},
		ImageUrl: pgtype.Text{
			String: req.GetUrl(),
			Valid:  req.Url != nil,
		},
		ImageThumbnails: pgtype.Text{
			String: req.GetThumbnails(),
			Valid:  req.Thumbnails != nil,
		},
		ImageBlurHash: pgtype.Text{
			String: req.GetBlurHash(),
			Valid:  req.BlurHash != nil,
		},
		ImageHeight: pgtype.Int4{
			Int32: int32(req.GetHeight()),
			Valid: req.Height != nil,
		},
		ImageWidth: pgtype.Int4{
			Int32: int32(req.GetWidth()),
			Valid: req.Width != nil,
		},
	}

	data, err := server.gojo.UpdateAnimeImage(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to update anime character", err)
	}

	res := &apbv1.UpdateAnimeImageResponse{
		AnimeImage: &apbv1.ImageResponse{
			ID:         data.ID,
			Host:       data.ImageHost,
			Url:        data.ImageUrl,
			Thumbnails: data.ImageThumbnails,
			BlurHash:   data.ImageBlurHash,
			Height:     uint32(data.ImageHeight),
			Width:      uint32(data.ImageWidth),
			CreatedAt:  timestamppb.New(data.CreatedAt),
		},
	}

	return res, nil
}

func validateUpdateAnimeImageRequest(req *apbv1.UpdateAnimeImageRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetImageID()); err != nil {
		violations = append(violations, shv1.FieldViolation("imageID", err))
	}

	if req.Host != nil {
		if err := utils.ValidateURL(req.GetHost(), ""); err != nil {
			violations = append(violations, shv1.FieldViolation("host", err))
		}
	}

	if req.Url != nil {
		if err := utils.ValidateImage(req.GetUrl()); err != nil {
			violations = append(violations, shv1.FieldViolation("url", err))
		}
	}

	if req.Thumbnails != nil {
		if err := utils.ValidateImage(req.GetThumbnails()); err != nil {
			violations = append(violations, shv1.FieldViolation("thumbnails", err))
		}
	}

	if req.BlurHash != nil {
		if err := utils.ValidateString(req.GetBlurHash(), 10, 50); err != nil {
			violations = append(violations, shv1.FieldViolation("blurHash", err))
		}
	}

	return violations
}
