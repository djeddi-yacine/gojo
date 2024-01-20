package asapiv1

import (
	"context"
	"errors"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSerieLink(ctx context.Context, req *aspbv1.CreateAnimeSerieLinkRequest) (*aspbv1.CreateAnimeSerieLinkResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie link")
	}

	if violations := validateCreateAnimeSerieLinkRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSerieLinkTxParams{
		AnimeID: req.GetAnimeID(),
		CreateAnimeLinkParams: db.CreateAnimeLinkParams{
			OfficialWebsite: req.GetAnimeLinks().GetOfficialWebsite(),
			WikipediaUrl:    req.GetAnimeLinks().GetWikipediaUrl(),
			CrunchyrollUrl:  req.GetAnimeLinks().GetCrunchyrollUrl(),
			SocialMedia:     req.GetAnimeLinks().GetSocialMedia(),
		},
	}

	data, err := server.gojo.CreateAnimeSerieLinkTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime serie link", err)
	}

	res := &aspbv1.CreateAnimeSerieLinkResponse{
		AnimeSerie: convertAnimeSerie(data.AnimeSerie),
		AnimeLinks: av1.ConvertAnimeLink(data.AnimeLink),
	}
	return res, nil
}

func validateCreateAnimeSerieLinkRequest(req *aspbv1.CreateAnimeSerieLinkRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("ID", err))
	}

	if req.AnimeLinks != nil {
		if err := utils.ValidateURL(req.GetAnimeLinks().GetOfficialWebsite(), ""); err != nil {
			violations = append(violations, shv1.FieldViolation("officialWebsite", err))
		}

		if err := utils.ValidateURL(req.GetAnimeLinks().GetCrunchyrollUrl(), "crunchyroll"); err != nil {
			violations = append(violations, shv1.FieldViolation("crunchyrollUrl", err))
		}

		if err := utils.ValidateURL(req.GetAnimeLinks().GetWikipediaUrl(), "wikipedia"); err != nil {
			violations = append(violations, shv1.FieldViolation("wikipediaUrl", err))
		}

		if len(req.GetAnimeLinks().GetSocialMedia()) > 0 {
			for _, v := range req.GetAnimeLinks().GetSocialMedia() {
				if err := utils.ValidateURL(v, ""); err != nil {
					violations = append(violations, shv1.FieldViolation("socialMedia", err))
				}
			}
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeLinks", errors.New("you need to send the AnimeLinks model")))
	}

	return violations
}
