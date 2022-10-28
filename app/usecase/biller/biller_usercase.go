package biller

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
	repo "github.com/vins7/emoney-service/app/adapter/client"
	db "github.com/vins7/emoney-service/app/adapter/db/e_money"
	"github.com/vins7/emoney-service/app/adapter/entity"
	"github.com/vins7/emoney-service/app/interface/model"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/e_money_service"
)

type BillerUsecase struct {
	repo repo.BillerRepo
	db   db.EMoneyRepo
}

func NewBillerUsecase(repo repo.BillerRepo, db db.EMoneyRepo) *BillerUsecase {
	return &BillerUsecase{
		repo: repo,
		db:   db,
	}
}

func (u *BillerUsecase) ListBiller() (interface{}, error) {
	res, err := u.repo.ListBiller()
	if err != nil {
		return nil, err
	}

	var req = &entity.ListBillerResponse{}
	if err := mapstructure.Decode(res, &req); err != nil {
		return nil, err
	}

	listBiller := []*proto.Biller{}
	for _, v := range req.ListBiller {
		listBiller = append(listBiller, &proto.Biller{
			ID:          strconv.Itoa(v.ID),
			Category:    v.Category,
			Product:     v.Product,
			Description: v.Description,
			Price:       strconv.Itoa(v.Price),
			Fee:         strconv.Itoa(v.Fee),
		})
	}

	return &proto.BillerResponse{
		ListBiller: listBiller,
	}, nil
}

func (u *BillerUsecase) DetailBiller(in interface{}) (interface{}, error) {
	var req *entity.BillerRequest

	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, err
	}

	res, err := u.repo.DetailBiller(req)
	if err != nil {
		return nil, err
	}

	return &proto.Biller{
		ID:          strconv.Itoa(res.DetailBiller.ID),
		Category:    res.DetailBiller.Category,
		Product:     res.DetailBiller.Product,
		Description: res.DetailBiller.Description,
		Price:       strconv.Itoa(res.DetailBiller.Price),
		Fee:         strconv.Itoa(res.DetailBiller.Fee),
	}, nil
}

func (u *BillerUsecase) GetBalance(in interface{}) (interface{}, error) {
	var req *model.GetBalance

	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, err
	}

	res, err := u.db.GetBalance(req)
	if err != nil {
		return nil, err
	}

	return &proto.GetBalanceResponse{
		UserId:  res.UserId,
		NoKartu: res.NoKartu,
		Balance: strconv.Itoa(*res.Balance),
	}, nil
}

func (u *BillerUsecase) CreateEMoney(in interface{}) (interface{}, error) {
	var req *model.CreateEmoney

	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, err
	}

	res, err := u.db.CreateEMoney(req)
	if err != nil {
		return nil, err
	}

	return &proto.CreateEMoneyResponse{
		UserId:  res.UserId,
		NoKartu: res.NoKartu,
	}, nil
}

func (u *BillerUsecase) TransactionHistory(in interface{}) (interface{}, error) {
	var req *model.GetTrxHist

	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, err
	}

	res, err := u.db.TransactionHistory(req)
	if err != nil {
		return nil, err
	}

	list := []*proto.TrxHist{}
	for _, v := range res {
		list = append(list, &proto.TrxHist{
			ID:          v.Id,
			UserId:      v.UserId,
			NoKartu:     v.NoKartu,
			CreatedDate: v.CreatedDate,
			UpdateDate:  v.UpdateDate,
			Setor:       v.Setor,
			Tarik:       v.Tarik,
			Balance:     v.Balance,
		})
	}

	return &proto.GetTrxHistResponse{
		ListTrxHist: list,
	}, nil
}
