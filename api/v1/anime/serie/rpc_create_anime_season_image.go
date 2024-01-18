package asapiv1

import (
	"context"
	"errors"
	"fmt"

	aapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSeasonImage(ctx context.Context, req *aspbv1.CreateAnimeSeasonImageRequest) (*aspbv1.CreateAnimeSeasonImageResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime season images")
	}

	if violations := validateCreateAnimeSeasonImageRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSeasonImageTxParams{
		SeasonID: req.GetSeasonID(),
	}

	if req.SeasonPosters != nil {
		arg.SeasonPosters = make([]db.CreateAnimeImageParams, len(req.GetSeasonPosters()))
		for i, v := range req.GetSeasonPosters() {
			arg.SeasonPosters[i].ImageHost = v.Host
			arg.SeasonPosters[i].ImageUrl = v.Url
			arg.SeasonPosters[i].ImageThumbnails = v.Thumbnails
			arg.SeasonPosters[i].ImageBlurhash = v.Blurhash
			arg.SeasonPosters[i].ImageHeight = int32(v.Height)
			arg.SeasonPosters[i].ImageWidth = int32(v.Width)
		}
	}

	data, err := server.gojo.CreateAnimeSeasonImageTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime season images", err)
	}

	res := &aspbv1.CreateAnimeSeasonImageResponse{
		AnimeSeason:   convertAnimeSeason(data.AnimeSeason),
		SeasonPosters: aapiv1.ConvertAnimeImages(data.AnimePosters),
	}
	return res, nil
}

func validateCreateAnimeSeasonImageRequest(req *aspbv1.CreateAnimeSeasonImageRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if req.SeasonPosters != nil {
		if len(req.GetSeasonPosters()) > 0 {
			for i, v := range req.GetSeasonPosters() {
				if err := utils.ValidateURL(v.Host, ""); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonPosters > host at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.Url, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonPosters > url at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.Thumbnails, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonPosters > thumbnails at index [%d]", i), err))
				}
				if err := utils.ValidateInt(int64(v.Height + 1)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonPosters > Height at index [%d]", i), err))
				}
				if err := utils.ValidateInt(int64(v.Width + 1)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonPosters > Width at index [%d]", i), err))
				}
			}
		}
	} else {
		violations = append(violations, shv1.FieldViolation("seasonPosters", errors.New("you need to send the posters in AnimeImages model")))
	}

	return violations
}
