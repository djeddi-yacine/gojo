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

func (server *AnimeMovieServer) CreateAnimeMovieTitle(ctx context.Context, req *ampbv1.CreateAnimeMovieTitleRequest) (*ampbv1.CreateAnimeMovieTitleResponse, error) {
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

	var DBF []db.CreateAnimeMovieOfficialTitleParams
	if req.AnimeTitles.Official != nil {
		DBF = make([]db.CreateAnimeMovieOfficialTitleParams, len(req.AnimeTitles.GetOfficial()))
		for i, t := range req.AnimeTitles.GetOfficial() {
			DBF[i].AnimeID = req.AnimeID
			DBF[i].TitleText = t
		}
	}

	var DBS []db.CreateAnimeMovieShortTitleParams
	if req.AnimeTitles.Short != nil {
		DBS = make([]db.CreateAnimeMovieShortTitleParams, len(req.AnimeTitles.GetShort()))
		for i, t := range req.AnimeTitles.GetShort() {
			DBS[i].AnimeID = req.AnimeID
			DBS[i].TitleText = t
		}
	}

	var DBT []db.CreateAnimeMovieOtherTitleParams
	if req.AnimeTitles.Other != nil {
		DBT = make([]db.CreateAnimeMovieOtherTitleParams, len(req.AnimeTitles.GetOther()))
		for i, t := range req.AnimeTitles.GetOther() {
			DBT[i].AnimeID = req.AnimeID
			DBT[i].TitleText = t
		}
	}

	arg := db.CreateAnimeMovieTitleTxParams{
		AnimeID:             req.GetAnimeID(),
		AnimeOfficialTitles: DBF,
		AnimeShortTitles:    DBS,
		AnimeOtherTitles:    DBT,
	}

	data, err := server.gojo.CreateAnimeMovieTitleTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime movie title", err)
	}

	var officials []*ampbv1.AnimeMovieTitle
	if len(data.AnimeOfficialTitles) > 0 {
		officials = make([]*ampbv1.AnimeMovieTitle, len(data.AnimeOfficialTitles))
		for i, t := range data.AnimeOfficialTitles {
			officials[i] = &ampbv1.AnimeMovieTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	var shorts []*ampbv1.AnimeMovieTitle
	if len(data.AnimeShortTitles) > 0 {
		shorts = make([]*ampbv1.AnimeMovieTitle, len(data.AnimeShortTitles))
		for i, t := range data.AnimeShortTitles {
			shorts[i] = &ampbv1.AnimeMovieTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	var others []*ampbv1.AnimeMovieTitle
	if len(data.AnimeOtherTitles) > 0 {
		shorts = make([]*ampbv1.AnimeMovieTitle, len(data.AnimeOtherTitles))
		for i, t := range data.AnimeOtherTitles {
			shorts[i] = &ampbv1.AnimeMovieTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	res := &ampbv1.CreateAnimeMovieTitleResponse{
		AnimeMovie: convertAnimeMovie(data.AnimeMovie),
		AnimeTitles: &ampbv1.AnimeMovieTitleResponse{
			Official: officials,
			Short:    shorts,
			Other:    others,
		},
	}

	return res, nil
}

func validateCreateAnimeMovieTitleRequest(req *ampbv1.CreateAnimeMovieTitleRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if req.AnimeTitles != nil {
		if req.AnimeTitles.Official != nil {
			if len(req.AnimeTitles.GetOfficial()) > 0 {
				for i, t := range req.AnimeTitles.GetOfficial() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTitles > official > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("animeTitles > official", errors.New("you need to send the official titles in animeTitles model")))
		}

		if req.AnimeTitles.Short != nil {
			if len(req.AnimeTitles.GetShort()) > 0 {
				for i, t := range req.AnimeTitles.GetShort() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
						violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTitles > short > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shv1.FieldViolation("animeTitles > short", errors.New("you need to send the short titles in animeTitles model")))
		}

		if req.AnimeTitles.Other != nil {
			if len(req.AnimeTitles.GetOther()) > 0 {
				for i, t := range req.AnimeTitles.GetOther() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
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
