package amapiv1

import (
	"context"
	"errors"
	"fmt"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) CreateAnimeMovieTitles(ctx context.Context, req *ampbv1.CreateAnimeMovieTitlesRequest) (*ampbv1.CreateAnimeMovieTitlesResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie title")
	}

	if violations := validateCreateAnimeMovieTitleRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieTitlesTxParams{
		AnimeID: req.GetAnimeID(),
	}

	if req.AnimeTitles.Official != nil {
		arg.AnimeOfficialTitles = make([]db.CreateAnimeMovieOfficialTitleParams, len(req.AnimeTitles.GetOfficial()))
		for i, v := range req.AnimeTitles.GetOfficial() {
			arg.AnimeOfficialTitles[i].AnimeID = req.AnimeID
			arg.AnimeOfficialTitles[i].TitleText = v
		}
	}

	if req.AnimeTitles.Short != nil {
		arg.AnimeShortTitles = make([]db.CreateAnimeMovieShortTitleParams, len(req.AnimeTitles.GetShort()))
		for i, v := range req.AnimeTitles.GetShort() {
			arg.AnimeShortTitles[i].AnimeID = req.AnimeID
			arg.AnimeShortTitles[i].TitleText = v
		}
	}

	if req.AnimeTitles.Other != nil {
		arg.AnimeOtherTitles = make([]db.CreateAnimeMovieOtherTitleParams, len(req.AnimeTitles.GetOther()))
		for i, v := range req.AnimeTitles.GetOther() {
			arg.AnimeOtherTitles[i].AnimeID = req.AnimeID
			arg.AnimeOtherTitles[i].TitleText = v
		}
	}

	data, err := server.gojo.CreateAnimeMovieTitlesTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime movie title", err)
	}

	var titles []string

	res := &ampbv1.CreateAnimeMovieTitlesResponse{
		AnimeTitles: &ampbv1.AnimeMovieTitleResponse{},
	}

	if len(data.AnimeOfficialTitles) > 0 {
		res.AnimeTitles.Official = make([]*ampbv1.AnimeMovieTitle, len(data.AnimeOfficialTitles))
		for i, v := range data.AnimeOfficialTitles {
			res.AnimeTitles.Official[i] = &ampbv1.AnimeMovieTitle{
				ID:        v.ID,
				TitleText: v.TitleText,
				CreatedAt: timestamppb.New(v.CreatedAt),
			}
			titles = append(titles, v.TitleText)
		}
	}

	if len(data.AnimeShortTitles) > 0 {
		res.AnimeTitles.Short = make([]*ampbv1.AnimeMovieTitle, len(data.AnimeShortTitles))
		for i, v := range data.AnimeShortTitles {
			res.AnimeTitles.Short[i] = &ampbv1.AnimeMovieTitle{
				ID:        v.ID,
				TitleText: v.TitleText,
				CreatedAt: timestamppb.New(v.CreatedAt),
			}
			titles = append(titles, v.TitleText)
		}
	}

	if len(data.AnimeOtherTitles) > 0 {
		res.AnimeTitles.Other = make([]*ampbv1.AnimeMovieTitle, len(data.AnimeOtherTitles))
		for i, v := range data.AnimeOtherTitles {
			res.AnimeTitles.Other[i] = &ampbv1.AnimeMovieTitle{
				ID:        v.ID,
				TitleText: v.TitleText,
				CreatedAt: timestamppb.New(v.CreatedAt),
			}
			titles = append(titles, v.TitleText)
		}
	}

	server.meilisearch.AddDocuments(&utils.Document{
		ID:     req.GetAnimeID(),
		Titles: utils.RemoveDuplicatesTitles(titles),
	})

	return res, nil
}

func validateCreateAnimeMovieTitleRequest(req *ampbv1.CreateAnimeMovieTitlesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if req.AnimeTitles != nil {
		if req.AnimeTitles.Official != nil {
			if len(req.AnimeTitles.GetOfficial()) > 0 {
				for i, v := range req.AnimeTitles.GetOfficial() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTitles > official > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("animeTitles > official", errors.New("you need to send the official titles in animeTitles model")))
		}

		if req.AnimeTitles.Short != nil {
			if len(req.AnimeTitles.GetShort()) > 0 {
				for i, v := range req.AnimeTitles.GetShort() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTitles > short > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("animeTitles > short", errors.New("you need to send the short titles in animeTitles model")))
		}

		if req.AnimeTitles.Other != nil {
			if len(req.AnimeTitles.GetOther()) > 0 {
				for i, v := range req.AnimeTitles.GetOther() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTitles > other > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("animeTitles > other", errors.New("you need to send the other titles in animeTitles model")))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeTitles", errors.New("you need to send the AnimeTitles model")))
	}

	return violations
}
