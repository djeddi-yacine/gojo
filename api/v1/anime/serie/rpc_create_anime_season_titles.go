package asapiv1

import (
	"context"
	"errors"
	"fmt"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) CreateAnimeSeasonTitles(ctx context.Context, req *aspbv1.CreateAnimeSeasonTitlesRequest) (*aspbv1.CreateAnimeSeasonTitlesResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime season title")
	}

	if violations := validateCreateAnimeSeasonTitleRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSeasonTitlesTxParams{
		SeasonID: req.GetSeasonID(),
	}

	if req.SeasonTitles.Official != nil {
		arg.AnimeOfficialTitles = make([]db.CreateAnimeSeasonOfficialTitleParams, len(req.SeasonTitles.GetOfficial()))
		for i, v := range req.SeasonTitles.GetOfficial() {
			arg.AnimeOfficialTitles[i].SeasonID = req.SeasonID
			arg.AnimeOfficialTitles[i].TitleText = v
		}
	}

	if req.SeasonTitles.Short != nil {
		arg.AnimeShortTitles = make([]db.CreateAnimeSeasonShortTitleParams, len(req.SeasonTitles.GetShort()))
		for i, v := range req.SeasonTitles.GetShort() {
			arg.AnimeShortTitles[i].SeasonID = req.SeasonID
			arg.AnimeShortTitles[i].TitleText = v
		}
	}

	if req.SeasonTitles.Other != nil {
		arg.AnimeOtherTitles = make([]db.CreateAnimeSeasonOtherTitleParams, len(req.SeasonTitles.GetOther()))
		for i, v := range req.SeasonTitles.GetOther() {
			arg.AnimeOtherTitles[i].SeasonID = req.SeasonID
			arg.AnimeOtherTitles[i].TitleText = v
		}
	}

	data, err := server.gojo.CreateAnimeSeasonTitlesTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime season title", err)
	}

	var titles []string
	res := &aspbv1.CreateAnimeSeasonTitlesResponse{
		SeasonTitles: &aspbv1.AnimeSeasonTitleResponse{},
	}

	if len(data.AnimeOfficialTitles) > 0 {
		res.SeasonTitles.Official = make([]*aspbv1.AnimeSeasonTitle, len(data.AnimeOfficialTitles))
		for i, v := range data.AnimeOfficialTitles {
			res.SeasonTitles.Official[i] = &aspbv1.AnimeSeasonTitle{
				ID:        v.ID,
				TitleText: v.TitleText,
				CreatedAt: timestamppb.New(v.CreatedAt),
			}
			titles = append(titles, v.TitleText)
		}
	}

	if len(data.AnimeShortTitles) > 0 {
		res.SeasonTitles.Short = make([]*aspbv1.AnimeSeasonTitle, len(data.AnimeShortTitles))
		for i, v := range data.AnimeShortTitles {
			res.SeasonTitles.Short[i] = &aspbv1.AnimeSeasonTitle{
				ID:        v.ID,
				TitleText: v.TitleText,
				CreatedAt: timestamppb.New(v.CreatedAt),
			}
			titles = append(titles, v.TitleText)
		}
	}

	if len(data.AnimeOtherTitles) > 0 {
		res.SeasonTitles.Other = make([]*aspbv1.AnimeSeasonTitle, len(data.AnimeOtherTitles))
		for i, v := range data.AnimeOtherTitles {
			res.SeasonTitles.Other[i] = &aspbv1.AnimeSeasonTitle{
				ID:        v.ID,
				TitleText: v.TitleText,
				CreatedAt: timestamppb.New(v.CreatedAt),
			}
			titles = append(titles, v.TitleText)
		}
	}

	server.meilisearch.AddDocuments(&utils.Document{
		ID:     req.GetSeasonID(),
		Titles: utils.RemoveDuplicatesTitles(titles),
	})

	return res, nil
}

func validateCreateAnimeSeasonTitleRequest(req *aspbv1.CreateAnimeSeasonTitlesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if req.SeasonTitles != nil {
		if req.SeasonTitles.Official != nil {
			if len(req.SeasonTitles.GetOfficial()) > 0 {
				for i, v := range req.SeasonTitles.GetOfficial() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonTitles > official > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("seasonTitles > official", errors.New("you need to send the official titles in seasonTitles model")))
		}

		if req.SeasonTitles.Short != nil {
			if len(req.SeasonTitles.GetShort()) > 0 {
				for i, v := range req.SeasonTitles.GetShort() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonTitles > short > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("seasonTitles > short", errors.New("you need to send the short titles in seasonTitles model")))
		}

		if req.SeasonTitles.Other != nil {
			if len(req.SeasonTitles.GetOther()) > 0 {
				for i, v := range req.SeasonTitles.GetOther() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonTitles > other > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("seasonTitles > other", errors.New("you need to send the other titles in seasonTitles model")))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("seasonTitles", errors.New("you need to send the seasonTitles model")))
	}

	return violations
}
