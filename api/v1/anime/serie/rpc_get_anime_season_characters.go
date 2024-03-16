package asapiv1

import (
	"context"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *AnimeSerieServer) GetAnimeSeasonCharacters(ctx context.Context, req *aspbv1.GetAnimeSeasonCharactersRequest) (*aspbv1.GetAnimeSeasonCharactersResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetAnimeSeasonCharactersRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := ping.CacheKey{
		ID:     req.SeasonID,
		Target: ping.AnimeSeason,
	}

	var characters []int64
	if err = server.ping.Handle(ctx, cache.Characters(), &characters, func() error {
		characters, err = server.gojo.ListAnimeSeasonCharacters(ctx, req.GetSeasonID())
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res := &aspbv1.GetAnimeSeasonCharactersResponse{}

	data, err := server.gojo.ListAnimeCharacetrsTx(ctx, characters)
	if err != nil {
		if dberr := db.ErrorDB(err); dberr != nil {
			if dberr.Code == pgerrcode.CaseNotFound {
				return res, nil
			}
		}

		return nil, shv1.ApiError("cannot get anime season characters", err)
	}

	res.SeasonCharacters = make([]*apbv1.AnimeCharacter, len(data))
	for i, v := range data {
		res.SeasonCharacters[i] = &apbv1.AnimeCharacter{
			Character: av1.ConvertAnimeCharacter(v.Character),
			Actors:    shv1.ConvertActors(v.Actor),
		}
	}

	return res, nil
}

func validateGetAnimeSeasonCharactersRequest(req *aspbv1.GetAnimeSeasonCharactersRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	return violations
}
