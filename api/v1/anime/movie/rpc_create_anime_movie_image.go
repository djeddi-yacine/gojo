package amapiv1

import (
	"context"
	"errors"
	"fmt"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) CreateAnimeMovieImage(ctx context.Context, req *ampbv1.CreateAnimeMovieImageRequest) (*ampbv1.CreateAnimeMovieImageResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie Image")
	}

	if violations := validateCreateAnimeMovieImageRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieImageTxParams{
		AnimeID: req.GetAnimeID(),
	}

	if req.AnimeImages.Posters != nil {
		arg.AnimePosters = make([]db.CreateAnimeImageParams, len(req.GetAnimeImages().GetPosters()))
		for i, v := range req.AnimeImages.GetPosters() {
			arg.AnimePosters[i].ImageHost = v.Host
			arg.AnimePosters[i].ImageUrl = v.Url
			arg.AnimePosters[i].ImageThumbnails = v.Thumbnails
			arg.AnimePosters[i].ImageBlurHash = v.BlurHash
			arg.AnimePosters[i].ImageHeight = int32(v.Height)
			arg.AnimePosters[i].ImageWidth = int32(v.Width)
		}
	}

	if req.AnimeImages.Backdrops != nil {
		arg.AnimeBackdrops = make([]db.CreateAnimeImageParams, len(req.GetAnimeImages().GetBackdrops()))
		for i, v := range req.AnimeImages.GetBackdrops() {
			arg.AnimeBackdrops[i].ImageHost = v.Host
			arg.AnimeBackdrops[i].ImageUrl = v.Url
			arg.AnimeBackdrops[i].ImageThumbnails = v.Thumbnails
			arg.AnimeBackdrops[i].ImageBlurHash = v.BlurHash
			arg.AnimeBackdrops[i].ImageHeight = int32(v.Height)
			arg.AnimeBackdrops[i].ImageWidth = int32(v.Width)
		}
	}

	if req.AnimeImages.Backdrops != nil {
		arg.AnimeLogos = make([]db.CreateAnimeImageParams, len(req.GetAnimeImages().GetLogos()))
		for i, v := range req.AnimeImages.GetLogos() {
			arg.AnimeLogos[i].ImageHost = v.Host
			arg.AnimeLogos[i].ImageUrl = v.Url
			arg.AnimeLogos[i].ImageThumbnails = v.Thumbnails
			arg.AnimeLogos[i].ImageBlurHash = v.BlurHash
			arg.AnimeLogos[i].ImageHeight = int32(v.Height)
			arg.AnimeLogos[i].ImageWidth = int32(v.Width)
		}
	}

	data, err := server.gojo.CreateAnimeMovieImageTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime movie Image", err)
	}

	res := &ampbv1.CreateAnimeMovieImageResponse{
		AnimeMovie: server.convertAnimeMovie(data.AnimeMovie),
		AnimeImages: &apbv1.AnimeImageResponse{
			Posters:   av1.ConvertAnimeImages(data.AnimePosters),
			Backdrops: av1.ConvertAnimeImages(data.AnimeBackdrops),
			Logos:     av1.ConvertAnimeImages(data.AnimeLogos),
		},
	}
	return res, nil
}

func validateCreateAnimeMovieImageRequest(req *ampbv1.CreateAnimeMovieImageRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if req.AnimeImages != nil {
		if req.AnimeImages.Posters == nil && req.AnimeImages.Backdrops == nil && req.AnimeImages.Logos == nil {
			violations = append(violations, shv1.FieldViolation("animeImages > posters,backdrops,logos", errors.New("you need to send one of [posters;backdrops;logos] in AnimeImages model")))
		}

		if len(req.AnimeImages.GetPosters()) > 0 {
			for i, v := range req.AnimeImages.GetPosters() {
				if err := utils.ValidateURL(v.Host, ""); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > posters > host at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.Url, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > posters > url at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.BlurHash, 1, 50); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > blurHash at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.Thumbnails, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > posters > thumbnails at index [%d]", i), err))
				}
				if err := utils.ValidateInt(int64(v.Height + 1)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > posters > Height at index [%d]", i), err))
				}
				if err := utils.ValidateInt(int64(v.Width + 1)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > posters > Width at index [%d]", i), err))
				}
			}
		}

		if len(req.AnimeImages.GetBackdrops()) > 0 {
			for i, v := range req.AnimeImages.GetBackdrops() {
				if err := utils.ValidateURL(v.Host, ""); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > backdrops > host at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.Url, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > backdrops > url at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.BlurHash, 1, 50); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > blurHash at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.Thumbnails, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > backdrops > thumbnails at index [%d]", i), err))
				}
				if err := utils.ValidateInt(int64(v.Height + 1)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > backdrops > Height at index [%d]", i), err))
				}
				if err := utils.ValidateInt(int64(v.Width + 1)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > backdrops > Width at index [%d]", i), err))
				}
			}
		}

		if len(req.AnimeImages.GetLogos()) > 0 {
			for i, v := range req.AnimeImages.GetLogos() {
				if err := utils.ValidateURL(v.Host, ""); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > logos > host at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.Url, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > logos > url at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.BlurHash, 1, 50); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > blurHash at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.Thumbnails, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > logos > thumbnails at index [%d]", i), err))
				}
				if err := utils.ValidateInt(int64(v.Height + 1)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > logos > Height at index [%d]", i), err))
				}
				if err := utils.ValidateInt(int64(v.Width + 1)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > logos > Width at index [%d]", i), err))
				}
			}
		}
	} else {
		violations = append(violations, shv1.FieldViolation("animeImages", errors.New("you need to send the AnimeImages model")))
	}

	return violations
}
