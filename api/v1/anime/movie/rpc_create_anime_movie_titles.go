package amapiv1

import (
	"context"
	"errors"
	"fmt"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
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

	if req.AnimeTitles.Officials != nil {
		arg.AnimeOfficialTitles = make([]db.CreateAnimeMovieOfficialTitleParams, len(req.AnimeTitles.GetOfficials()))
		for i, v := range req.AnimeTitles.GetOfficials() {
			arg.AnimeOfficialTitles[i].AnimeID = req.AnimeID
			arg.AnimeOfficialTitles[i].TitleText = v
		}
	}

	if req.AnimeTitles.Shorts != nil {
		arg.AnimeShortTitles = make([]db.CreateAnimeMovieShortTitleParams, len(req.AnimeTitles.GetShorts()))
		for i, v := range req.AnimeTitles.GetShorts() {
			arg.AnimeShortTitles[i].AnimeID = req.AnimeID
			arg.AnimeShortTitles[i].TitleText = v
		}
	}

	if req.AnimeTitles.Others != nil {
		arg.AnimeOtherTitles = make([]db.CreateAnimeMovieOtherTitleParams, len(req.AnimeTitles.GetOthers()))
		for i, v := range req.AnimeTitles.GetOthers() {
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
		AnimeTitles: &apbv1.AnimeTitlesResponse{},
	}

	if len(data.AnimeOfficialTitles) > 0 {
		res.AnimeTitles.Officials = make([]*apbv1.AnimeTitle, len(data.AnimeOfficialTitles))
		for i, v := range data.AnimeOfficialTitles {
			res.AnimeTitles.Officials[i] = &apbv1.AnimeTitle{
				ID:        v.ID,
				TitleText: v.TitleText,
				CreatedAt: timestamppb.New(v.CreatedAt),
			}
			titles = append(titles, v.TitleText)
		}
	}

	if len(data.AnimeShortTitles) > 0 {
		res.AnimeTitles.Shorts = make([]*apbv1.AnimeTitle, len(data.AnimeShortTitles))
		for i, v := range data.AnimeShortTitles {
			res.AnimeTitles.Shorts[i] = &apbv1.AnimeTitle{
				ID:        v.ID,
				TitleText: v.TitleText,
				CreatedAt: timestamppb.New(v.CreatedAt),
			}
			titles = append(titles, v.TitleText)
		}
	}

	if len(data.AnimeOtherTitles) > 0 {
		res.AnimeTitles.Others = make([]*apbv1.AnimeTitle, len(data.AnimeOtherTitles))
		for i, v := range data.AnimeOtherTitles {
			res.AnimeTitles.Others[i] = &apbv1.AnimeTitle{
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
		if req.AnimeTitles.Officials != nil {
			if len(req.AnimeTitles.GetOfficials()) > 0 {
				for i, v := range req.AnimeTitles.GetOfficials() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTitles > officials > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("animeTitles > officials", errors.New("you need to send the official titles in animeTitles model")))
		}

		if req.AnimeTitles.Shorts != nil {
			if len(req.AnimeTitles.GetShorts()) > 0 {
				for i, v := range req.AnimeTitles.GetShorts() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTitles > shorts > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("animeTitles > shorts", errors.New("you need to send the short titles in animeTitles model")))
		}

		if req.AnimeTitles.Others != nil {
			if len(req.AnimeTitles.GetOthers()) > 0 {
				for i, v := range req.AnimeTitles.GetOthers() {
					if err := utils.ValidateString(v, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTitles > others > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("animeTitles > others", errors.New("you need to send the other titles in animeTitles model")))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeTitles", errors.New("you need to send the AnimeTitles model")))
	}

	return violations
}
