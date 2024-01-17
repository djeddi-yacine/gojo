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

	var DBF []db.CreateAnimeSeasonOfficialTitleParams
	if req.SeasonTitles.Official != nil {
		DBF = make([]db.CreateAnimeSeasonOfficialTitleParams, len(req.SeasonTitles.GetOfficial()))
		for i, t := range req.SeasonTitles.GetOfficial() {
			DBF[i].SeasonID = req.SeasonID
			DBF[i].TitleText = t
		}
	}

	var DBS []db.CreateAnimeSeasonShortTitleParams
	if req.SeasonTitles.Short != nil {
		DBS = make([]db.CreateAnimeSeasonShortTitleParams, len(req.SeasonTitles.GetShort()))
		for i, t := range req.SeasonTitles.GetShort() {
			DBS[i].SeasonID = req.SeasonID
			DBS[i].TitleText = t
		}
	}

	var DBT []db.CreateAnimeSeasonOtherTitleParams
	if req.SeasonTitles.Other != nil {
		DBT = make([]db.CreateAnimeSeasonOtherTitleParams, len(req.SeasonTitles.GetOther()))
		for i, t := range req.SeasonTitles.GetOther() {
			DBT[i].SeasonID = req.SeasonID
			DBT[i].TitleText = t
		}
	}

	arg := db.CreateAnimeSeasonTitlesTxParams{
		SeasonID:            req.GetSeasonID(),
		AnimeOfficialTitles: DBF,
		AnimeShortTitles:    DBS,
		AnimeOtherTitles:    DBT,
	}

	data, err := server.gojo.CreateAnimeSeasonTitlesTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime season title", err)
	}

	var titles []string

	var officials []*aspbv1.AnimeSeasonTitle
	if len(data.AnimeOfficialTitles) > 0 {
		officials = make([]*aspbv1.AnimeSeasonTitle, len(data.AnimeOfficialTitles))
		for i, t := range data.AnimeOfficialTitles {
			officials[i] = &aspbv1.AnimeSeasonTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
			titles = append(titles, t.TitleText)
		}
	}

	var shorts []*aspbv1.AnimeSeasonTitle
	if len(data.AnimeShortTitles) > 0 {
		shorts = make([]*aspbv1.AnimeSeasonTitle, len(data.AnimeShortTitles))
		for i, t := range data.AnimeShortTitles {
			shorts[i] = &aspbv1.AnimeSeasonTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
			titles = append(titles, t.TitleText)
		}
	}

	var others []*aspbv1.AnimeSeasonTitle
	if len(data.AnimeOtherTitles) > 0 {
		shorts = make([]*aspbv1.AnimeSeasonTitle, len(data.AnimeOtherTitles))
		for i, t := range data.AnimeOtherTitles {
			shorts[i] = &aspbv1.AnimeSeasonTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
			titles = append(titles, t.TitleText)
		}
	}

	server.meilisearch.AddDocuments(&utils.Document{
		ID:     req.GetSeasonID(),
		Titles: utils.RemoveDuplicatesTitles(titles),
	})

	res := &aspbv1.CreateAnimeSeasonTitlesResponse{
		SeasonTitles: &aspbv1.AnimeSeasonTitleResponse{
			Official: officials,
			Short:    shorts,
			Other:    others,
		},
	}

	return res, nil
}

func validateCreateAnimeSeasonTitleRequest(req *aspbv1.CreateAnimeSeasonTitlesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if req.SeasonTitles != nil {
		if req.SeasonTitles.Official != nil {
			if len(req.SeasonTitles.GetOfficial()) > 0 {
				for i, t := range req.SeasonTitles.GetOfficial() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonTitles > official > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("seasonTitles > official", errors.New("you need to send the official titles in seasonTitles model")))
		}

		if req.SeasonTitles.Short != nil {
			if len(req.SeasonTitles.GetShort()) > 0 {
				for i, t := range req.SeasonTitles.GetShort() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonTitles > short > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("seasonTitles > short", errors.New("you need to send the short titles in seasonTitles model")))
		}

		if req.SeasonTitles.Other != nil {
			if len(req.SeasonTitles.GetOther()) > 0 {
				for i, t := range req.SeasonTitles.GetOther() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
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
