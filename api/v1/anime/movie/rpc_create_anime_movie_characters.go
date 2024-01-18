package amapiv1

import (
	"context"
	"errors"
	"fmt"

	aapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	ashpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ashpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) CreateAnimeMovieCharacters(ctx context.Context, req *ampbv1.CreateAnimeMovieCharactersRequest) (*ampbv1.CreateAnimeMovieCharactersResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie characters")
	}

	if violations := validateCreateAnimeMovieCharactersRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieCharactersTxParams{
		AnimeID: req.GetAnimeID(),
	}

	arg.AnimeCharacterActorsTxParams = make([]db.AnimeCharacterActorsTxParams, len(req.GetAnimeCharacters()))
	for i, v := range req.GetAnimeCharacters() {
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

	data, err := server.gojo.CreateAnimeMovieCharactersTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime movie characters", err)
	}

	res := &ampbv1.CreateAnimeMovieCharactersResponse{
		AnimeMovie: convertAnimeMovie(data.AnimeMovie),
	}

	res.AnimeCharacters = make([]*ashpbv1.AnimeCharacterResponse, len(data.Characters))
	for i, v := range data.Characters {
		res.AnimeCharacters[i] = aapiv1.ConvertAnimeCharacter(v.AnimeCharacter)
		res.AnimeCharacters[i].ActorsID = v.ActorsIDs
	}

	return res, nil
}

func validateCreateAnimeMovieCharactersRequest(req *ampbv1.CreateAnimeMovieCharactersRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if len(req.AnimeCharacters) > 0 {
		for i, v := range req.GetAnimeCharacters() {
			if err := utils.ValidateString(v.GetFullName(), 2, 100); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeCharacters > fullName at index [%d]", i), err))
			}

			if err := utils.ValidateString(v.GetAbout(), 2, 10000); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeCharacters > about at index [%d]", i), err))
			}

			if err := utils.ValidateString(v.GetImageBlurHash(), 10, 50); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeCharacters > imageBlurHash at index [%d]", i), err))
			}

			if err := utils.ValidateURL(v.GetImage(), ""); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeCharacters > image at index [%d]", i), err))
			}

			if len(v.GetActorsID()) <= 0 {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeCharacters > actorsID at index [%d]", i), fmt.Errorf("put one actorID at least")))
			}
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeCharacters", errors.New("you need to send the at least one of AnimeCharacters model")))
	}

	return violations
}
