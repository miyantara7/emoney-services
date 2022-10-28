package e_money

import (
	"github.com/vins7/emoney-service/app/adapter/entity"
	"github.com/vins7/emoney-service/app/interface/model"
)

type EMoneyRepo interface {
	GetBalance(req *model.GetBalance) (*entity.EMoney, error)
	TransactionHistory(req *model.GetTrxHist) ([]*entity.TransactionHistory, error)
	CreateEMoney(req *model.CreateEmoney) (*entity.EMoney, error)
}
