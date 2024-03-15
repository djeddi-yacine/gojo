package nfapiv1

import (
	"context"
	"errors"
	"fmt"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) CreateActors(ctx context.Context, req *nfpbv1.CreateActorsRequest) (*nfpbv1.CreateActorsResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create new actors")
	}

	if violations := validateCreateActorsRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	actors := make([]db.CreateActorParams, len(req.GetActors()))
	for i, v := range req.GetActors() {
		actors[i] = db.CreateActorParams{
			FullName:      v.GetFullName(),
			Biography:     v.GetBiography(),
			Gender:        v.GetGender(),
			Born:          v.GetBorn().AsTime(),
			ImageUrl:      v.GetImage(),
			ImageBlurHash: v.GetImageBlurHash(),
		}
	}

	result, err := server.gojo.CreateActorsTx(ctx, actors)
	if err != nil {
		return nil, shv1.ApiError("failed to create new actors", err)
	}

	res := &nfpbv1.CreateActorsResponse{
		Actors: shv1.ConvertActors(result),
	}

	return res, nil
}

func validateCreateActorsRequest(req *nfpbv1.CreateActorsRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if len(req.GetActors()) > 0 {
		for i, v := range req.GetActors() {
			if err := utils.ValidateString(v.GetFullName(), 2, 100); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("fullName at index [%d]", i), err))
			}

			if err := utils.ValidateString(v.GetBiography(), 2, 5000); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("biography at index [%d]", i), err))
			}

			if err := utils.ValidateString(v.GetGender(), 2, 20); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("gender at index [%d]", i), err))
			}

			if err := utils.ValidateURL(v.GetImage(), ""); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("image at index [%d]", i), err))
			}

			if err := utils.ValidateString(v.GetImageBlurHash(), 2, 32); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("imageBlurhash at index [%d]", i), err))
			}
		}
	} else {
		violations = append(violations, shv1.FieldViolation("actors", errors.New("you need to send at least one ActorRequest model")))
	}

	return violations
}
