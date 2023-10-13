package api

import (
	"context"
	"fmt"
	"time"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateAnimeMovie(ctx context.Context, req *pb.CreateAnimeMovieRequest) (*pb.CreateAnimeMovieResponse, error) {
	if violations := validateCreateAnimeMovieRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	argAnimeTx := db.CreateAnimeMovieTxParams{
		CreateAnimeMovieParams: db.CreateAnimeMovieParams{
			OriginalTitle: req.AnimeMovie.GetOriginalTitle(),
			Aired:         req.AnimeMovie.GetAired().AsTime(),
			ReleaseYear:   req.AnimeMovie.GetReleaseYear(),
			Duration:      req.AnimeMovie.GetDuration(),
		},
		GenreIDs:  req.GetAnimeGenres().GetGenreID(),
		StudioIDs: req.GetAnimeStudios().GetStudioID(),
	}

	resultAnime, err := server.gojo.CreateAnimeMovieTx(ctx, argAnimeTx)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime movie : %s", err)
	}

	res := &pb.CreateAnimeMovieResponse{
		AnimeMovie: &pb.AnimeMovieRes{
			OriginalTitle: req.AnimeMovie.GetOriginalTitle(),
			Aired:         timestamppb.New(resultAnime.AnimeMovie.Aired),
			Premiered:     fmt.Sprint(resultAnime.AnimeMovie.ReleaseYear),
			Duration:      resultAnime.AnimeMovie.Duration,
			CreatedAt:     timestamppb.New(resultAnime.AnimeMovie.CreatedAt),
		},
	}
	return res, nil
}

func validateCreateAnimeMovieRequest(req *pb.CreateAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.AnimeMovie.OriginalTitle != "" {
		if err := utils.ValidateString(req.GetAnimeMovie().GetOriginalTitle(), 1, 100); err != nil {
			violations = append(violations, fieldViolation("originalTitle", err))
		}
	}

	if req.AnimeMovie.Aired != nil {
		if err := utils.ValidateDate(req.GetAnimeMovie().GetAired().AsTime().Format(time.DateOnly)); err != nil {
			violations = append(violations, fieldViolation("aired", err))
		}
	}

	if req.AnimeMovie.ReleaseYear != 0 {
		if err := utils.ValidateYear(req.GetAnimeMovie().GetReleaseYear()); err != nil {
			violations = append(violations, fieldViolation("releaseYear", err))
		}
	}

	return violations
}
