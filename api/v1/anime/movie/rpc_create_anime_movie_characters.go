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

	dbCharacters := make([]db.AnimeCharacterActorsTxParams, len(req.GetAnimeCharacters()))

	for i, x := range req.GetAnimeCharacters() {
		dbCharacters[i].CreateAnimeCharacters = db.CreateAnimeCharacterParams{
			FullName:      x.GetFullName(),
			About:         x.GetAbout(),
			RolePlaying:   x.GetRolePlaying(),
			ImageUrl:      x.GetImage(),
			ImageBlurHash: x.GetImageBlurHash(),
			Pictures:      x.GetPictures(),
		}

		dbCharacters[i].ActorsIDs = x.GetActorsID()
	}

	data, err := server.gojo.CreateAnimeMovieCharactersTx(ctx, db.CreateAnimeMovieCharactersTxParams{
		AnimeID:                      req.GetAnimeID(),
		AnimeCharacterActorsTxParams: dbCharacters,
	})
	if err != nil {
		return nil, shv1.ApiError("failed to create anime movie characters", err)
	}

	res := &ampbv1.CreateAnimeMovieCharactersResponse{
		AnimeMovie: convertAnimeMovie(data.AnimeMovie),
	}

	res.AnimeCharacters = make([]*ashpbv1.AnimeCharacterResponse, len(data.Characters))
	for i, x := range data.Characters {
		res.AnimeCharacters[i] = aapiv1.ConvertAnimeCharacter(x.AnimeCharacter)
		res.AnimeCharacters[i].ActorsID = x.ActorsIDs
	}

	return res, nil
}

func validateCreateAnimeMovieCharactersRequest(req *ampbv1.CreateAnimeMovieCharactersRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if len(req.AnimeCharacters) > 0 {
		for i, x := range req.GetAnimeCharacters() {
			if err := utils.ValidateString(x.GetFullName(), 2, 100); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeCharacters > fullName at index [%d]", i), err))
			}

			if err := utils.ValidateString(x.GetAbout(), 2, 10000); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeCharacters > about at index [%d]", i), err))
			}

			if err := utils.ValidateString(x.GetImageBlurHash(), 10, 50); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeCharacters > imageBlurHash at index [%d]", i), err))
			}

			if err := utils.ValidateURL(x.GetImage(), ""); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeCharacters > image at index [%d]", i), err))
			}

			if len(x.GetActorsID()) <= 0 {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeCharacters > actorsID at index [%d]", i), fmt.Errorf("put one actorID at least")))
			}
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeCharacters", errors.New("you need to send the at least one of AnimeCharacters model")))
	}

	return violations
}
