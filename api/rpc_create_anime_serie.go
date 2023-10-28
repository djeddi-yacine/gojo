package api

import (
	"context"
	"time"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateAnimeSerie(ctx context.Context, req *aspb.CreateAnimeSerieRequest) (*aspb.CreateAnimeSerieResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie")
	}

	if violations := validateCreateAnimeSerieRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateAnimeSerieParams{
		OriginalTitle: req.AnimeSerie.GetOriginalTitle(),
		Aired:         req.AnimeSerie.GetAired().AsTime(),
		ReleaseYear:   req.AnimeSerie.GetReleaseYear(),
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
			OriginalTitle: req.AnimeSerie.GetOriginalTitle(),
			Aired:         timestamppb.New(anime.Aired),
			ReleaseYear:   anime.ReleaseYear,
			Duration:      durationpb.New(anime.Duration),
			CreatedAt:     timestamppb.New(anime.CreatedAt),
		},
	}
	return res, nil
}

func validateCreateAnimeSerieRequest(req *aspb.CreateAnimeSerieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateString(req.GetAnimeSerie().GetOriginalTitle(), 2, 500); err != nil {
		violations = append(violations, fieldViolation("originalTitle", err))
	}

	if err := utils.ValidateDate(req.GetAnimeSerie().GetAired().AsTime().Format(time.DateOnly)); err != nil {
		violations = append(violations, fieldViolation("aired", err))
	}

	if err := utils.ValidateYear(req.GetAnimeSerie().GetReleaseYear()); err != nil {
		violations = append(violations, fieldViolation("releaseYear", err))
	}

	if err := utils.ValidateDuration(req.GetAnimeSerie().GetDuration().AsDuration().String()); err != nil {
		violations = append(violations, fieldViolation("duration", err))
	}

	return violations
}
