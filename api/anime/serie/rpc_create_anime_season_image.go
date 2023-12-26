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

func (server *AnimeSerieServer) CreateAnimeSeasonImage(ctx context.Context, req *aspb.CreateAnimeSeasonImageRequest) (*aspb.CreateAnimeSeasonImageResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime season images")
	}

	if violations := validateCreateAnimeSeasonImageRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	var DBP []db.CreateAnimeImageParams
	if req.SeasonPosters != nil {
		DBP = make([]db.CreateAnimeImageParams, len(req.GetSeasonPosters()))
		for i, p := range req.GetSeasonPosters() {
			DBP[i].ImageHost = p.Host
			DBP[i].ImageUrl = p.Url
			DBP[i].ImageThumbnails = p.Thumbnails
			DBP[i].ImageBlurhash = p.Blurhash
			DBP[i].ImageHeight = int32(p.Height)
			DBP[i].ImageWidth = int32(p.Width)
		}
	}

	arg := db.CreateAnimeSeasonImageTxParams{
		SeasonID:      req.GetSeasonID(),
		SeasonPosters: DBP,
	}

	data, err := server.gojo.CreateAnimeSeasonImageTx(ctx, arg)
	if err != nil {
		return nil, shared.ApiError("failed to create anime season images", err)
	}

	res := &aspb.CreateAnimeSeasonImageResponse{
		AnimeSeason:   shared.ConvertAnimeSeason(data.AnimeSeason),
		SeasonPosters: shared.ConvertAnimeImages(data.AnimePosters),
	}
	return res, nil
}

func validateCreateAnimeSeasonImageRequest(req *aspb.CreateAnimeSeasonImageRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shared.FieldViolation("seasonID", err))
	}

	if req.SeasonPosters != nil {
		if len(req.GetSeasonPosters()) > 0 {
			for i, l := range req.GetSeasonPosters() {
				if err := utils.ValidateURL(l.Host, ""); err != nil {
					violations = append(violations, shared.FieldViolation(fmt.Sprintf("seasonPosters > host at index [%d]", i), err))
				}
				if err := utils.ValidateString(l.Url, 1, 200); err != nil {
					violations = append(violations, shared.FieldViolation(fmt.Sprintf("seasonPosters > url at index [%d]", i), err))
				}
				if err := utils.ValidateString(l.Thumbnails, 1, 200); err != nil {
					violations = append(violations, shared.FieldViolation(fmt.Sprintf("seasonPosters > thumbnails at index [%d]", i), err))
				}
				if err := utils.ValidateInt(int64(l.Height + 1)); err != nil {
					violations = append(violations, shared.FieldViolation(fmt.Sprintf("seasonPosters > Height at index [%d]", i), err))
				}
				if err := utils.ValidateInt(int64(l.Width + 1)); err != nil {
					violations = append(violations, shared.FieldViolation(fmt.Sprintf("seasonPosters > Width at index [%d]", i), err))
				}
			}
		}
	} else {
		violations = append(violations, shared.FieldViolation("seasonPosters", errors.New("you need to send the posters in AnimeImages model")))
	}

	return violations
}
