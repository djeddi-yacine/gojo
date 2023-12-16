package animeMovie

import (
	"context"
	"errors"
	"fmt"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) CreateAnimeMovieTitle(ctx context.Context, req *ampb.CreateAnimeMovieTitleRequest) (*ampb.CreateAnimeMovieTitleResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie title")
	}

	if violations := validateCreateAnimeMovieTitleRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
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
		return nil, shared.DatabaseError("failed to create anime movie title", err)
	}

	var officials []*ampb.AnimeMovieTitle
	if len(data.AnimeOfficialTitles) > 0 {
		officials = make([]*ampb.AnimeMovieTitle, len(data.AnimeOfficialTitles))
		for i, t := range data.AnimeOfficialTitles {
			officials[i] = &ampb.AnimeMovieTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	var shorts []*ampb.AnimeMovieTitle
	if len(data.AnimeShortTitles) > 0 {
		shorts = make([]*ampb.AnimeMovieTitle, len(data.AnimeShortTitles))
		for i, t := range data.AnimeShortTitles {
			shorts[i] = &ampb.AnimeMovieTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	var others []*ampb.AnimeMovieTitle
	if len(data.AnimeOtherTitles) > 0 {
		shorts = make([]*ampb.AnimeMovieTitle, len(data.AnimeOtherTitles))
		for i, t := range data.AnimeOtherTitles {
			shorts[i] = &ampb.AnimeMovieTitle{
				ID:        t.ID,
				TitleText: t.TitleText,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	res := &ampb.CreateAnimeMovieTitleResponse{
		AnimeMovie: shared.ConvertAnimeMovie(data.AnimeMovie),
		AnimeTitles: &ampb.AnimeMovieTitleResponse{
			Official: officials,
			Short:    shorts,
			Other:    others,
		},
	}

	return res, nil
}

func validateCreateAnimeMovieTitleRequest(req *ampb.CreateAnimeMovieTitleRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("animeID", err))
	}

	if req.AnimeTitles != nil {
		if req.AnimeTitles.Official != nil {
			if len(req.AnimeTitles.GetOfficial()) > 0 {
				for i, t := range req.AnimeTitles.GetOfficial() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeTitles > official > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shared.FieldViolation("animeTitles > official", errors.New("you need to send the official titles in animeTitles model")))
		}

		if req.AnimeTitles.Short != nil {
			if len(req.AnimeTitles.GetShort()) > 0 {
				for i, t := range req.AnimeTitles.GetShort() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeTitles > short > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shared.FieldViolation("animeTitles > short", errors.New("you need to send the short titles in animeTitles model")))
		}

		if req.AnimeTitles.Other != nil {
			if len(req.AnimeTitles.GetOther()) > 0 {
				for i, t := range req.AnimeTitles.GetOther() {
					if err := utils.ValidateString(t, 1, 150); err != nil {
						violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeTitles > other > title at index [%d]", i), err))
					}
				}
			}
		} else {
			violations = append(violations, shared.FieldViolation("animeTitles > other", errors.New("you need to send the other titles in animeTitles model")))
		}

	} else {
		violations = append(violations, shared.FieldViolation("animeTitles", errors.New("you need to send the AnimeTitles model")))
	}

	return violations
}
