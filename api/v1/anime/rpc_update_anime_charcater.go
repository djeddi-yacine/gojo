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

func (server *AnimeServer) UpdateAnimeCharacter(ctx context.Context, req *apbv1.UpdateAnimeCharacterRequest) (*apbv1.UpdateAnimeCharacterResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update anime character")
	}

	if violations := validateUpdateAnimeCharacterRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	data, err := server.gojo.UpdateAnimeCharacterTx(ctx, db.UpdateAnimeCharacterTxParams{
		AnimeCharacter: db.UpdateAnimeCharacterParams{
			ID: req.GetCharacterID(),
			FullName: pgtype.Text{
				String: req.GetFullName(),
				Valid:  req.FullName != nil,
			},
			About: pgtype.Text{
				String: req.GetAbout(),
				Valid:  req.About != nil,
			},
			RolePlaying: pgtype.Text{
				String: req.GetRolePlaying(),
				Valid:  req.RolePlaying != nil,
			},
			ImageUrl: pgtype.Text{
				String: req.GetImage(),
				Valid:  req.Image != nil,
			},
			ImageBlurHash: pgtype.Text{
				String: req.GetImageBlurHash(),
				Valid:  req.ImageBlurHash != nil,
			},
			Pictures: req.GetPictures(),
		},
		ActorsIDs: req.GetActorsID(),
	})
	if err != nil {
		return nil, shv1.ApiError("failed to update anime character", err)
	}

	return &apbv1.UpdateAnimeCharacterResponse{
		AnimeCharacter: &apbv1.AnimeCharacterResponse{
			ID:            data.AnimeCharacter.ID,
			FullName:      data.AnimeCharacter.FullName,
			About:         data.AnimeCharacter.About,
			RolePlaying:   data.AnimeCharacter.RolePlaying,
			Image:         data.AnimeCharacter.ImageUrl,
			ImageBlurHash: data.AnimeCharacter.ImageBlurHash,
			Pictures:      data.AnimeCharacter.Pictures,
			ActorsID:      data.ActorsIDs,
			CreatedAt:     timestamppb.New(data.AnimeCharacter.CreatedAt),
		},
	}, nil
}

func validateUpdateAnimeCharacterRequest(req *apbv1.UpdateAnimeCharacterRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetCharacterID()); err != nil {
		violations = append(violations, shv1.FieldViolation("characterID", err))
	}

	if req.FullName != nil {
		if err := utils.ValidateString(req.GetFullName(), 2, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("fullName", err))
		}
	}

	if req.About != nil {
		if err := utils.ValidateString(req.GetAbout(), 2, 10000); err != nil {
			violations = append(violations, shv1.FieldViolation("about", err))
		}
	}

	if req.Image != nil {
		if err := utils.ValidateURL(req.GetImage(), ""); err != nil {
			violations = append(violations, shv1.FieldViolation("image", err))
		}
	}

	if req.ImageBlurHash != nil {
		if err := utils.ValidateString(req.GetImageBlurHash(), 10, 50); err != nil {
			violations = append(violations, shv1.FieldViolation("imageBlurHash", err))
		}
	}

	if req.Pictures != nil {
		for _, v := range req.Pictures {
			if err := utils.ValidateURL(v, ""); err != nil {
				violations = append(violations, shv1.FieldViolation("pictures", err))
			}
		}
	}

	if req.ActorsID != nil {
		if len(req.GetActorsID()) <= 0 {
			violations = append(violations, shv1.FieldViolation("actorsID", errors.New("put one actorID at least")))
		}
	}

	return violations
}
