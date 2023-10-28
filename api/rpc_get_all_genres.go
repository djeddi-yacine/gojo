package api

import (
	"context"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAllGenres(ctx context.Context, req *nfpb.GetAllGenresRequest) (*nfpb.GetAllGenresResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all genres")
	}

	violations := validateGetAllGenresRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListGenresParams{
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	}
	DBgenres, err := server.gojo.ListGenres(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to list the genres : %s", err)
	}

	var PBgenres []*nfpb.Genre
	for _, g := range DBgenres {
		PBgenres = append(PBgenres, ConvertGenre(g))
	}

	res := &nfpb.GetAllGenresResponse{
		Genres: PBgenres,
	}
	return res, nil
}

func validateGetAllGenresRequest(req *nfpb.GetAllGenresRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, fieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, fieldViolation("pageSize", err))
	}

	return violations
}
