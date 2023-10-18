package api

import (
	"context"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAllStudios(ctx context.Context, req *pb.GetAllStudiosRequest) (*pb.GetAllStudiosResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all Studios")
	}

	violations := validateGetAllStudiosRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListStudiosParams{
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	}
	DBStudios, err := server.gojo.ListStudios(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to list the Studios : %s", err)
	}

	var PBStudios []*pb.Studio
	for _, g := range DBStudios {
		PBStudios = append(PBStudios, ConvertStudio(g))
	}

	res := &pb.GetAllStudiosResponse{
		Studios: PBStudios,
	}
	return res, nil
}

func validateGetAllStudiosRequest(req *pb.GetAllStudiosRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetPageNumber()); err != nil {
		violations = append(violations, fieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("pageSize", err))
	}

	return violations
}
