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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) CreateAnimeSeasonTitle(ctx context.Context, req *aspb.CreateAnimeSeasonTitleRequest) (*aspb.CreateAnimeSeasonTitleResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime season title")
	}

	if violations := validateCreateAnimeSeasonTitleRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
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

	arg := db.CreateAnimeSeasonTitleTxParams{
		SeasonID:            req.GetSeasonID(),
		AnimeOfficialTitles: DBF,
		AnimeShortTitles:    DBS,
		AnimeOtherTitles:    DBT,
	}

	data, err := server.gojo.CreateAnimeSeasonTitleTx(ctx, arg)
	if err != nil {
		return nil, shared.ApiError("failed to create anime season title", err)
	}

	var officials []*aspb.AnimeSeasonTitle
	if len(data.AnimeOfficialTitles) > 0 {
		officials = make([]*aspb.AnimeSeasonTitle, len(data.AnimeOfficialTitles))
		for i, t := range data.AnimeOfficialTitles {
			officials[i] = &aspb.AnimeSeasonTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	var shorts []*aspb.AnimeSeasonTitle
	if len(data.AnimeShortTitles) > 0 {
		shorts = make([]*aspb.AnimeSeasonTitle, len(data.AnimeShortTitles))
		for i, t := range data.AnimeShortTitles {
			shorts[i] = &aspb.AnimeSeasonTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	var others []*aspb.AnimeSeasonTitle
	if len(data.AnimeOtherTitles) > 0 {
		shorts = make([]*aspb.AnimeSeasonTitle, len(data.AnimeOtherTitles))
		for i, t := range data.AnimeOtherTitles {
			shorts[i] = &aspb.AnimeSeasonTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	res := &aspb.CreateAnimeSeasonTitleResponse{
		AnimeSeason: shared.ConvertAnimeSeason(data.AnimeSeason),
		SeasonTitles: &aspb.AnimeSeasonTitleResponse{
			Official: officials,
			Short:    shorts,
			Other:    others,
		},
	}

	return res, nil
}

func validateCreateAnimeSeasonTitleRequest(req *aspb.CreateAnimeSeasonTitleRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shared.FieldViolation("seasonID", err))
	}

	if req.SeasonTitles != nil {
		if req.SeasonTitles.Official != nil {
			if len(req.SeasonTitles.GetOfficial()) > 0 {
				for i, t := range req.SeasonTitles.GetOfficial() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("seasonTitles > official > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shared.FieldViolation("seasonTitles > official", errors.New("you need to send the official titles in seasonTitles model")))
		}

		if req.SeasonTitles.Short != nil {
			if len(req.SeasonTitles.GetShort()) > 0 {
				for i, t := range req.SeasonTitles.GetShort() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("seasonTitles > short > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shared.FieldViolation("seasonTitles > short", errors.New("you need to send the short titles in seasonTitles model")))
		}

		if req.SeasonTitles.Other != nil {
			if len(req.SeasonTitles.GetOther()) > 0 {
				for i, t := range req.SeasonTitles.GetOther() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("seasonTitles > other > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shared.FieldViolation("seasonTitles > other", errors.New("you need to send the other titles in seasonTitles model")))
		}

	} else {
		violations = append(violations, shared.FieldViolation("seasonTitles", errors.New("you need to send the seasonTitles model")))
	}

	return violations
}
