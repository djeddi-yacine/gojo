package api

import (
	"context"
	"time"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateAnimeMovie(ctx context.Context, req *pb.CreateAnimeMovieRequest) (*pb.CreateAnimeMovieResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie")
	}

	if violations := validateCreateAnimeMovieRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieParams{
		OriginalTitle: req.AnimeMovie.GetOriginalTitle(),
		Aired:         req.AnimeMovie.GetAired().AsTime(),
		ReleaseYear:   req.AnimeMovie.GetReleaseYear(),
		Duration:      req.AnimeMovie.GetDuration().AsDuration(),
	}

	anime, err := server.gojo.CreateAnimeMovie(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime movie : %s", err)
	}

	res := &pb.CreateAnimeMovieResponse{
		AnimeMovie: &pb.AnimeMovieResponse{
			ID:            anime.ID,
			OriginalTitle: req.AnimeMovie.GetOriginalTitle(),
			Aired:         timestamppb.New(anime.Aired),
			ReleaseYear:   anime.ReleaseYear,
			Duration:      durationpb.New(anime.Duration),
			CreatedAt:     timestamppb.New(anime.CreatedAt),
		},
	}
	return res, nil
}

func validateCreateAnimeMovieRequest(req *pb.CreateAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateString(req.GetAnimeMovie().GetOriginalTitle(), 2, 500); err != nil {
		violations = append(violations, fieldViolation("originalTitle", err))
	}

	if err := utils.ValidateDate(req.GetAnimeMovie().GetAired().AsTime().Format(time.DateOnly)); err != nil {
		violations = append(violations, fieldViolation("aired", err))
	}

	if err := utils.ValidateYear(req.GetAnimeMovie().GetReleaseYear()); err != nil {
		violations = append(violations, fieldViolation("releaseYear", err))
	}

	return violations
}
