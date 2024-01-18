package asapiv1

import (
	"context"
	"errors"
	"fmt"

	aapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ashpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ashpb"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSerieImage(ctx context.Context, req *aspbv1.CreateAnimeSerieImageRequest) (*aspbv1.CreateAnimeSerieImageResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie mages")
	}

	if violations := validateCreateAnimeSerieImageRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSerieImageTxParams{
		AnimeID: req.GetAnimeID(),
	}

	if req.AnimeImages.Posters != nil {
		arg.AnimePosters = make([]db.CreateAnimeImageParams, len(req.GetAnimeImages().GetPosters()))
		for i, v := range req.AnimeImages.GetPosters() {
			arg.AnimePosters[i].ImageHost = v.Host
			arg.AnimePosters[i].ImageUrl = v.Url
			arg.AnimePosters[i].ImageThumbnails = v.Thumbnails
			arg.AnimePosters[i].ImageBlurhash = v.Blurhash
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
			arg.AnimeBackdrops[i].ImageBlurhash = v.Blurhash
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
			arg.AnimeLogos[i].ImageBlurhash = v.Blurhash
			arg.AnimeLogos[i].ImageHeight = int32(v.Height)
			arg.AnimeLogos[i].ImageWidth = int32(v.Width)
		}
	}

	data, err := server.gojo.CreateAnimeSerieImageTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime serie images", err)
	}

	res := &aspbv1.CreateAnimeSerieImageResponse{
		AnimeSerie: convertAnimeSerie(data.AnimeSerie),
		AnimeImages: &ashpbv1.AnimeImageResponse{
			Posters:   aapiv1.ConvertAnimeImages(data.AnimePosters),
			Backdrops: aapiv1.ConvertAnimeImages(data.AnimeBackdrops),
			Logos:     aapiv1.ConvertAnimeImages(data.AnimeLogos),
		},
	}
	return res, nil
}

func validateCreateAnimeSerieImageRequest(req *aspbv1.CreateAnimeSerieImageRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("ID", err))
	}

	if req.AnimeImages != nil {
		if req.AnimeImages.Posters == nil && req.AnimeImages.Backdrops == nil && req.AnimeImages.Logos == nil {
			violations = append(violations, shv1.FieldViolation("animeImages > posters;backdrops;logos", errors.New("you need to send one of [posters;backdrops;logos] in AnimeImages model")))
		}

		if len(req.AnimeImages.GetPosters()) > 0 {
			for i, v := range req.AnimeImages.GetPosters() {
				if err := utils.ValidateURL(v.Host, ""); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > posters > host at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.Url, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeImages > posters > url at index [%d]", i), err))
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
