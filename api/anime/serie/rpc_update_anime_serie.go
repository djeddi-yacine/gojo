package animeSerie

import (
	"context"
	"fmt"
	"time"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) UpdateAnimeSerie(ctx context.Context, req *aspb.UpdateAnimeSerieRequest) (*aspb.UpdateAnimeSerieResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update anime serie")
	}

	if violations := validateUpdateAnimeSerieRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.UpdateAnimeSerieParams{
		ID: req.ID,
		OriginalTitle: pgtype.Text{
			String: req.GetOriginalTitle(),
			Valid:  req.OriginalTitle != nil,
		},
		Aired: pgtype.Timestamptz{
			Time:  req.GetAired().AsTime(),
			Valid: req.Aired != nil,
		},
		ReleaseYear: pgtype.Int4{
			Int32: req.GetReleaseYear(),
			Valid: req.ReleaseYear != nil,
		},
		Duration: pgtype.Interval{
			Microseconds: req.GetDuration().AsDuration().Microseconds(),
			Valid:        req.Duration != nil,
		},
	}

	anime, err := server.gojo.UpdateAnimeSerie(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to update anime serie : %s", err)
	}

	res := &aspb.UpdateAnimeSerieResponse{
		AnimeSerie: &aspb.AnimeSerieResponse{
			ID:            anime.ID,
			OriginalTitle: req.GetOriginalTitle(),
			Aired:         timestamppb.New(anime.Aired),
			ReleaseYear:   anime.ReleaseYear,
			Duration:      durationpb.New(anime.Duration),
			CreatedAt:     timestamppb.New(anime.CreatedAt),
		},
	}
	return res, nil
}

func validateUpdateAnimeSerieRequest(req *aspb.UpdateAnimeSerieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.ID); err != nil {
		violations = append(violations, shared.FieldViolation("ID", err))
	}

	if req.OriginalTitle != nil {
		if err := utils.ValidateString(req.GetOriginalTitle(), 2, 500); err != nil {
			violations = append(violations, shared.FieldViolation("originalTitle", err))
		}
	}

	if req.Aired != nil {
		if err := utils.ValidateDate(req.GetAired().AsTime().Format(time.DateOnly)); err != nil {
			violations = append(violations, shared.FieldViolation("aired", err))
		}
	}

	if req.ReleaseYear != nil {
		if err := utils.ValidateYear(req.GetReleaseYear()); err != nil {
			violations = append(violations, shared.FieldViolation("releaseYear", err))
		}
	}

	if req.Duration != nil {
		fmt.Println(req.GetDuration().String())
		if err := utils.ValidateDuration(req.GetDuration().AsDuration().String()); err != nil {
			violations = append(violations, shared.FieldViolation("duration", err))
		}
	}

	return violations
}
