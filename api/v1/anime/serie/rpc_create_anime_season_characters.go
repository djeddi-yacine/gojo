package asapiv1

import (
	"context"
	"errors"
	"fmt"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ashpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ashpb"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSeasonCharacters(ctx context.Context, req *aspbv1.CreateAnimeSeasonCharactersRequest) (*aspbv1.CreateAnimeSeasonCharactersResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime season characters")
	}

	if violations := validateCreateAnimeSeasonCharactersRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSeasonCharactersTxParams{
		SeasonID: req.GetSeasonID(),
	}

	arg.AnimeCharacterActorsTxParams = make([]db.AnimeCharacterActorsTxParams, len(req.GetSeasonCharacters()))
	for i, v := range req.GetSeasonCharacters() {
		arg.AnimeCharacterActorsTxParams[i].CreateAnimeCharacters = db.CreateAnimeCharacterParams{
			FullName:      v.GetFullName(),
			About:         v.GetAbout(),
			RolePlaying:   v.GetRolePlaying(),
			ImageUrl:      v.GetImage(),
			ImageBlurHash: v.GetImageBlurHash(),
			Pictures:      v.GetPictures(),
		}

		arg.AnimeCharacterActorsTxParams[i].ActorsIDs = v.GetActorsID()
	}

	data, err := server.gojo.CreateAnimeSeasonCharactersTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime season characters", err)
	}

	res := &aspbv1.CreateAnimeSeasonCharactersResponse{
		AnimeSeason: convertAnimeSeason(data.AnimeSeason),
	}

	res.SeasonCharacters = make([]*ashpbv1.AnimeCharacterResponse, len(data.Characters))
	for i, v := range data.Characters {
		res.SeasonCharacters[i] = av1.ConvertAnimeCharacter(v.AnimeCharacter)
		res.SeasonCharacters[i].ActorsID = v.ActorsIDs
	}

	return res, nil
}

func validateCreateAnimeSeasonCharactersRequest(req *aspbv1.CreateAnimeSeasonCharactersRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if len(req.GetSeasonCharacters()) > 0 {
		for i, v := range req.GetSeasonCharacters() {
			if err := utils.ValidateString(v.GetFullName(), 2, 100); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonCharacters > fullName at index [%d]", i), err))
			}

			if err := utils.ValidateString(v.GetAbout(), 2, 10000); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonCharacters > about at index [%d]", i), err))
			}

			if err := utils.ValidateString(v.GetImageBlurHash(), 10, 50); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonCharacters > imageBlurHash at index [%d]", i), err))
			}

			if err := utils.ValidateURL(v.GetImage(), ""); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonCharacters > image at index [%d]", i), err))
			}

			if len(v.GetActorsID()) <= 0 {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonCharacters > actorsID at index [%d]", i), fmt.Errorf("put one actorID at least")))
			}
		}

	} else {
		violations = append(violations, shv1.FieldViolation("seasonCharacters", errors.New("you need to send the at least one of AnimeCharacters model")))
	}

	return violations
}
