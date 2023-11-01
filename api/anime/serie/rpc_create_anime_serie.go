package animeSerie

import (
	"context"
	"time"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) CreateAnimeSerie(ctx context.Context, req *aspb.CreateAnimeSerieRequest) (*aspb.CreateAnimeSerieResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie")
	}

	if violations := validateCreateAnimeSerieRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSerieParams{
		OriginalTitle: req.AnimeSerie.GetOriginalTitle(),
		Aired:         req.AnimeSerie.GetAired().AsTime(),
		ReleaseYear:   req.AnimeSerie.GetReleaseYear(),
		Rating:        req.AnimeSerie.GetRating(),
		Duration:      req.AnimeSerie.GetDuration().AsDuration(),
	}

	anime, err := server.gojo.CreateAnimeSerie(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime serie : %s", err)
	}

	res := &aspb.CreateAnimeSerieResponse{
		AnimeSerie: &aspb.AnimeSerieResponse{
			ID:            anime.ID,
			OriginalTitle: anime.OriginalTitle,
			Aired:         timestamppb.New(anime.Aired),
			ReleaseYear:   anime.ReleaseYear,
			Rating:        anime.Rating,
			Duration:      durationpb.New(anime.Duration),
			CreatedAt:     timestamppb.New(anime.CreatedAt),
		},
	}
	return res, nil
}

func validateCreateAnimeSerieRequest(req *aspb.CreateAnimeSerieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateString(req.GetAnimeSerie().GetOriginalTitle(), 2, 500); err != nil {
		violations = append(violations, shared.FieldViolation("originalTitle", err))
	}

	if err := utils.ValidateDate(req.GetAnimeSerie().GetAired().AsTime().Format(time.DateOnly)); err != nil {
		violations = append(violations, shared.FieldViolation("aired", err))
	}

	if err := utils.ValidateYear(req.GetAnimeSerie().GetReleaseYear()); err != nil {
		violations = append(violations, shared.FieldViolation("releaseYear", err))
	}

	if err := utils.ValidateString(req.GetAnimeSerie().GetRating(), 2, 30); err != nil {
		violations = append(violations, shared.FieldViolation("rating", err))
	}

	if err := utils.ValidateDuration(req.GetAnimeSerie().GetDuration().AsDuration().String()); err != nil {
		violations = append(violations, shared.FieldViolation("duration", err))
	}

	return violations
}