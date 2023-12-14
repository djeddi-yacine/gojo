package animeSerie

import (
	"context"
	"errors"
	"fmt"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSeasonTrailer(ctx context.Context, req *aspb.CreateAnimeSeasonTrailerRequest) (*aspb.CreateAnimeSeasonTrailerResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime season trailer")
	}

	if violations := validateCreateAnimeSeasonTrailerRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	var DBT []db.CreateAnimeTrailerParams
	if req.SeasonTrailers != nil {
		DBT = make([]db.CreateAnimeTrailerParams, len(req.GetSeasonTrailers()))
		for i, t := range req.GetSeasonTrailers() {
			DBT[i].IsOfficial = t.IsOfficial
			DBT[i].HostName = t.HostName
			DBT[i].HostKey = t.HostKey
		}
	}

	arg := db.CreateAnimeSeasonTrailerTxParams{
		SeasonID:             req.GetSeasonID(),
		SeasonTrailersParams: DBT,
	}

	data, err := server.gojo.CreateAnimeSeasonTrailerTx(ctx, arg)
	if err != nil {
		return nil, shared.DatabaseError("failed to create anime season trailer", err)
	}

	res := &aspb.CreateAnimeSeasonTrailerResponse{
		AnimeSeason:    shared.ConvertAnimeSerieSeason(data.AnimeSeason),
		SeasonTrailers: shared.ConvertAnimeTrailers(data.SeasonTrailers),
	}
	return res, nil
}

func validateCreateAnimeSeasonTrailerRequest(req *aspb.CreateAnimeSeasonTrailerRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shared.FieldViolation("ID", err))
	}

	if req.SeasonTrailers != nil {
		if len(req.GetSeasonTrailers()) > 0 {
			for i, t := range req.GetSeasonTrailers() {
				if err := utils.ValidateString(t.HostName, 1, 200); err != nil {
					violations = append(violations, shared.FieldViolation(fmt.Sprintf("SeasonTrailers > hostName at index [%d]", i), err))
				}
				if err := utils.ValidateString(t.HostKey, 1, 200); err != nil {
					violations = append(violations, shared.FieldViolation(fmt.Sprintf("SeasonTrailers > hostKey at index [%d]", i), err))
				}
			}
		}

	} else {
		violations = append(violations, shared.FieldViolation("SeasonTrailers", errors.New("you need to send the SeasonTrailers model")))
	}

	return violations
}
