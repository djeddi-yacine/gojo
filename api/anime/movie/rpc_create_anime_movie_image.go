package animeMovie

import (
	"context"
	"errors"
	"fmt"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/pb/shpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) CreateAnimeMovieImage(ctx context.Context, req *ampb.CreateAnimeMovieImageRequest) (*ampb.CreateAnimeMovieImageResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie Image")
	}

	if violations := validateCreateAnimeMovieImageRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	var DBP []db.CreateAnimeImageParams
	if req.AnimeImages.Posters != nil {
		DBP = make([]db.CreateAnimeImageParams, len(req.AnimeImages.GetPosters()))
		for i, p := range req.AnimeImages.GetPosters() {
			DBP[i].ImageHost = p.Host
			DBP[i].ImageUrl = p.Url
			DBP[i].ImageThumbnails = p.Thumbnails
			DBP[i].ImageBlurhash = p.Blurhash
			DBP[i].ImageHeight = int32(p.Height)
			DBP[i].ImageWidth = int32(p.Width)
		}
	}

	var DBB []db.CreateAnimeImageParams
	if req.AnimeImages.Backdrops != nil {
		DBB = make([]db.CreateAnimeImageParams, len(req.AnimeImages.GetBackdrops()))
		for i, p := range req.AnimeImages.GetBackdrops() {
			DBB[i].ImageHost = p.Host
			DBB[i].ImageUrl = p.Url
			DBB[i].ImageThumbnails = p.Thumbnails
			DBB[i].ImageBlurhash = p.Blurhash
			DBB[i].ImageHeight = int32(p.Height)
			DBB[i].ImageWidth = int32(p.Width)
		}
	}

	var DBL []db.CreateAnimeImageParams
	if req.AnimeImages.Backdrops != nil {
		DBL = make([]db.CreateAnimeImageParams, len(req.AnimeImages.GetLogos()))
		for i, p := range req.AnimeImages.GetLogos() {
			DBL[i].ImageHost = p.Host
			DBL[i].ImageUrl = p.Url
			DBL[i].ImageThumbnails = p.Thumbnails
			DBL[i].ImageBlurhash = p.Blurhash
			DBL[i].ImageHeight = int32(p.Height)
			DBL[i].ImageWidth = int32(p.Width)
		}
	}

	arg := db.CreateAnimeMovieImageTxParams{
		AnimeID:        req.GetAnimeID(),
		AnimePosters:   DBP,
		AnimeBackdrops: DBB,
		AnimeLogos:     DBL,
	}

	data, err := server.gojo.CreateAnimeMovieImageTx(ctx, arg)
	if err != nil {
		return nil, shared.DatabaseError("failed to create anime movie Image", err)
	}

	res := &ampb.CreateAnimeMovieImageResponse{
		AnimeMovie: shared.ConvertAnimeMovie(data.AnimeMovie),
		AnimeImages: &shpb.AnimeImageResponse{
			Posters:   shared.ConvertAnimeImages(data.AnimePosters),
			Backdrops: shared.ConvertAnimeImages(data.AnimeBackdrops),
			Logos:     shared.ConvertAnimeImages(data.AnimeLogos),
		},
	}
	return res, nil
}

func validateCreateAnimeMovieImageRequest(req *ampb.CreateAnimeMovieImageRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("animeID", err))
	}

	if req.AnimeImages != nil {
		if req.AnimeImages.Posters != nil {
			if len(req.AnimeImages.GetPosters()) > 0 {
				for i, l := range req.AnimeImages.GetPosters() {
					if err := utils.ValidateURL(l.Host, ""); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > posters > host at index [%d]", i), err))
					}
					if err := utils.ValidateString(l.Url, 1, 200); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > posters > url at index [%d]", i), err))
					}
					if err := utils.ValidateString(l.Thumbnails, 1, 200); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > posters > thumbnails at index [%d]", i), err))
					}
					if err := utils.ValidateInt(int64(l.Height + 1)); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > posters > Height at index [%d]", i), err))
					}
					if err := utils.ValidateInt(int64(l.Width + 1)); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > posters > Width at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shared.FieldViolation("animeImages > posters", errors.New("you need to send the posters in AnimeImages model")))
		}

		if req.AnimeImages.Backdrops != nil {
			if len(req.AnimeImages.GetBackdrops()) > 0 {
				for i, l := range req.AnimeImages.GetBackdrops() {
					if err := utils.ValidateURL(l.Host, ""); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > backdrops > host at index [%d]", i), err))
					}
					if err := utils.ValidateString(l.Url, 1, 200); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > backdrops > url at index [%d]", i), err))
					}
					if err := utils.ValidateString(l.Thumbnails, 1, 200); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > backdrops > thumbnails at index [%d]", i), err))
					}
					if err := utils.ValidateInt(int64(l.Height + 1)); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > backdrops > Height at index [%d]", i), err))
					}
					if err := utils.ValidateInt(int64(l.Width + 1)); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > backdrops > Width at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shared.FieldViolation("animeImages > backdrops", errors.New("you need to send the posters in AnimeImages model")))
		}

		if req.AnimeImages.Logos != nil {
			if len(req.AnimeImages.GetLogos()) > 0 {
				for i, l := range req.AnimeImages.GetLogos() {
					if err := utils.ValidateURL(l.Host, ""); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > logos > host at index [%d]", i), err))
					}
					if err := utils.ValidateString(l.Url, 1, 200); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > logos > url at index [%d]", i), err))
					}
					if err := utils.ValidateString(l.Thumbnails, 1, 200); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > logos > thumbnails at index [%d]", i), err))
					}
					if err := utils.ValidateInt(int64(l.Height + 1)); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > logos > Height at index [%d]", i), err))
					}
					if err := utils.ValidateInt(int64(l.Width + 1)); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeImages > logos > Width at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shared.FieldViolation("animeImages > logos", errors.New("you need to send the posters in AnimeImages model")))
		}

	} else {
		violations = append(violations, shared.FieldViolation("animeImages", errors.New("you need to send the AnimeImages model")))
	}

	return violations
}
