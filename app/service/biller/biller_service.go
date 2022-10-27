package biller

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	ucBiller "github.com/vins7/emoney-service/app/usecase/biller"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/e_money_service"
)

type BillerService struct {
	uc ucBiller.BillerUsecase
}

func NewBillerService(uc ucBiller.BillerUsecase) *BillerService {
	return &BillerService{
		uc: uc,
	}
}

func (u *BillerService) ListBiller(ctx context.Context, req *empty.Empty) (*proto.BillerResponse, error) {
	_, err := u.uc.ListBiller(req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *BillerService) DetailBiller(ctx context.Context, req *empty.Empty) (*proto.Biller, error) {
	_, err := u.uc.ListBiller(req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
