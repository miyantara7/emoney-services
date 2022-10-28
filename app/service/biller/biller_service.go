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
	res, err := u.uc.ListBiller()
	if err != nil {
		return nil, err
	}

	return res.(*proto.BillerResponse), nil
}

func (u *BillerService) DetailBiller(ctx context.Context, req *proto.BillerRequest) (*proto.Biller, error) {
	res, err := u.uc.DetailBiller(req)
	if err != nil {
		return nil, err
	}

	return res.(*proto.Biller), nil
}

func (u *BillerService) GetBalance(ctx context.Context, req *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	res, err := u.uc.GetBalance(req)
	if err != nil {
		return nil, err
	}

	return res.(*proto.GetBalanceResponse), nil
}

func (u *BillerService) GetTrxHist(ctx context.Context, req *proto.GetTrxHistReq) (*proto.GetTrxHistResponse, error) {
	res, err := u.uc.TransactionHistory(req)
	if err != nil {
		return nil, err
	}

	return res.(*proto.GetTrxHistResponse), nil
}

func (u *BillerService) CreateEMoney(ctx context.Context, req *proto.CreateEMoneyRequest) (*proto.CreateEMoneyResponse, error) {
	res, err := u.uc.CreateEMoney(req)
	if err != nil {
		return nil, err
	}

	return res.(*proto.CreateEMoneyResponse), nil
}
