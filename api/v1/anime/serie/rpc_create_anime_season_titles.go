package asapiv1

import (
	"context"
	"errors"
	"fmt"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
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

	if req.SeasonTitles.Officials != nil {
		arg.AnimeOfficialTitles = make([]db.CreateAnimeSeasonOfficialTitleParams, len(req.SeasonTitles.GetOfficials()))
		for i, v := range req.SeasonTitles.GetOfficials() {
			arg.AnimeOfficialTitles[i].SeasonID = req.SeasonID
			arg.AnimeOfficialTitles[i].TitleText = v
		}
	}

	if req.SeasonTitles.Shorts != nil {
		arg.AnimeShortTitles = make([]db.CreateAnimeSeasonShortTitleParams, len(req.SeasonTitles.GetShorts()))
		for i, v := range req.SeasonTitles.GetShorts() {
			arg.AnimeShortTitles[i].SeasonID = req.SeasonID
			arg.AnimeShortTitles[i].TitleText = v
		}
	}

	if req.SeasonTitles.Others != nil {
		arg.AnimeOtherTitles = make([]db.CreateAnimeSeasonOtherTitleParams, len(req.SeasonTitles.GetOthers()))
		for i, v := range req.SeasonTitles.GetOthers() {
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
		SeasonTitles: &apbv1.AnimeTitlesResponse{},
	}

	if len(data.AnimeOfficialTitles) > 0 {
		res.SeasonTitles.Officials = make([]*apbv1.AnimeTitle, len(data.AnimeOfficialTitles))
		for i, v := range data.AnimeOfficialTitles {
			res.SeasonTitles.Officials[i] = &apbv1.AnimeTitle{
				ID:        v.ID,
				TitleText: v.TitleText,
				CreatedAt: timestamppb.New(v.CreatedAt),
			}
			titles = append(titles, v.TitleText)
		}
	}

	if len(data.AnimeShortTitles) > 0 {
		res.SeasonTitles.Shorts = make([]*apbv1.AnimeTitle, len(data.AnimeShortTitles))
		for i, v := range data.AnimeShortTitles {
			res.SeasonTitles.Shorts[i] = &apbv1.AnimeTitle{
				ID:        v.ID,
				TitleText: v.TitleText,
				CreatedAt: timestamppb.New(v.CreatedAt),
			}
			titles = append(titles, v.TitleText)
		}
	}

	if len(data.AnimeOtherTitles) > 0 {
		res.SeasonTitles.Others = make([]*apbv1.AnimeTitle, len(data.AnimeOtherTitles))
		for i, v := range data.AnimeOtherTitles {
			res.SeasonTitles.Others[i] = &apbv1.AnimeTitle{
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
		if req.SeasonTitles.Officials != nil {
			if len(req.SeasonTitles.GetOfficials()) > 0 {
				for i, v := range req.SeasonTitles.GetOfficials() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonTitles > officials > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("seasonTitles > officials", errors.New("you need to send the official titles in seasonTitles model")))
		}

		if req.SeasonTitles.Shorts != nil {
			if len(req.SeasonTitles.GetShorts()) > 0 {
				for i, v := range req.SeasonTitles.GetShorts() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonTitles > shorts > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("seasonTitles > shorts", errors.New("you need to send the short titles in seasonTitles model")))
		}

		if req.SeasonTitles.Others != nil {
			if len(req.SeasonTitles.GetOthers()) > 0 {
				for i, v := range req.SeasonTitles.GetOthers() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonTitles > others > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("seasonTitles > others", errors.New("you need to send the other titles in seasonTitles model")))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("seasonTitles", errors.New("you need to send the seasonTitles model")))
	}

	return violations
}
