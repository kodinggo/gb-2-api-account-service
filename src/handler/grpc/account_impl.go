package grpc

import (
	"context"

	pb "github.com/kodinggo/gb-2-api-account-service/pb/account"
	"github.com/kodinggo/gb-2-api-account-service/src/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountgRPCHandler struct {
	pb.UnimplementedAccountServiceServer
	usecase model.AccountUsecase
}

func NewAccountgRPCHandler(usecase model.AccountUsecase) pb.AccountServiceServer {
	return &AccountgRPCHandler{usecase: usecase}
}

func (a *AccountgRPCHandler) FindByID(ctx context.Context, req *pb.FindByIDsRequest) (*pb.Account, error) {
	account, err := a.usecase.FindByID(ctx, req.Ids[0])
	if err != nil {
		if err.Error() == "not found" {
			return nil, status.Errorf(codes.NotFound, "account with ID %d not found", req.Ids[0])
		}
		return nil, status.Errorf(codes.Internal, "error finding account: %v", err)
	}

	return &pb.Account{
		Id:         account.ID,
		Fullname:   account.Fullname,
		SortBio:    account.SortBio,
		Gender:     convertGenderToProto(account.Gender),
		PictureUrl: account.PictureUrl,
		Username:   account.Username,
		Email:      account.Email,
	}, nil

}

func convertGenderToProto(gender model.Gender) pb.Account_Gender {
	switch gender {
	case model.MALE:
		return pb.Account_MALE
	case model.FEMALE:
		return pb.Account_FEMALE
	case model.OTHERS:
		return pb.Account_OTHERS
	default:
		return pb.Account_OTHERS
	}
}
