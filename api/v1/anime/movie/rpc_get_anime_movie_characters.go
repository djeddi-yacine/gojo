package amapiv1

import (
	"context"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *AnimeMovieServer) GetAnimeMovieCharacters(ctx context.Context, req *ampbv1.GetAnimeMovieCharactersRequest) (*ampbv1.GetAnimeMovieCharactersResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetAnimeMovieCharactersRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := ping.CacheKey{
		ID:     req.AnimeID,
		Target: ping.AnimeMovie,
	}

	var characters []int64
	if err = server.ping.Handle(ctx, cache.Characters(), &characters, func() error {
		characters, err = server.gojo.ListAnimeMovieCharacters(ctx, req.GetAnimeID())
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res := &ampbv1.GetAnimeMovieCharactersResponse{}

	data, err := server.gojo.ListAnimeCharacetrsTx(ctx, characters)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shv1.ApiError("cannot get anime movie characters", err)
	}

	res.AnimeCharacters = make([]*apbv1.AnimeCharacter, len(data))
	for i, v := range data {
		res.AnimeCharacters[i] = &apbv1.AnimeCharacter{
			Character: av1.ConvertAnimeCharacter(v.Character),
			Actors:    shv1.ConvertActors(v.Actor),
		}
	}

	return res, nil
}

func validateGetAnimeMovieCharactersRequest(req *ampbv1.GetAnimeMovieCharactersRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	return violations
}
