package e_money

import (
	"errors"
	"strconv"
	"time"

	"github.com/vins7/emoney-service/app/adapter/entity"
	"github.com/vins7/emoney-service/app/interface/model"
	"github.com/vins7/emoney-service/app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type EMoneyDB struct {
	dbUser   *gorm.DB
	dbEMoney *gorm.DB
}

func NewEMoneyDB(db *gorm.DB, dbEMoney *gorm.DB) *EMoneyDB {
	return &EMoneyDB{
		dbUser:   db,
		dbEMoney: dbEMoney,
	}
}

func (u *EMoneyDB) GetBalance(req *model.GetBalance) (*entity.EMoney, error) {

	var data *entity.EMoney
	if err := u.dbEMoney.Debug().Where("no_kartu = ? and user_id = ?", req.NoKartu, req.UserId).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "no kartu not exist !")
		}
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return data, nil
}

func (u *EMoneyDB) DetailUser(req *model.DetailUserReq) (*entity.User, error) {

	data := &entity.User{}
	if err := u.dbUser.Debug().
		Joins("DataUser").
		Where("username = ? and DataUser.user_id = ?", req.Username, req.UserId).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user tidak ditemukan !")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return data, nil
}

func (u *EMoneyDB) CreateEMoney(req *model.CreateEmoney) (*entity.EMoney, error) {

	dataUser, err := u.DetailUser(&model.DetailUserReq{
		Username: req.UserName,
		UserId:   req.UserId,
	})
	if err != nil {
		return nil, err
	}

	data := &entity.EMoney{}
	err = u.dbEMoney.Transaction(func(tx *gorm.DB) error {

		saldo := 0
		if req.Saldo != "" {
			saldo, err = strconv.Atoi(req.Saldo)
			if err != nil {
				return status.Errorf(codes.Internal, err.Error())
			}
		}

		data = &entity.EMoney{
			NoKartu:    util.GenerateNoKartu(),
			UserId:     req.UserId,
			DataUserId: dataUser.Id,
			Balance:    &saldo,
		}

		if err := tx.Debug().Create(data).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

		if err := tx.Debug().Create(&entity.TransactionHistory{
			UserId:      req.UserId,
			NoKartu:     data.NoKartu,
			CreatedDate: time.Now().Format("2006-01-02 15:04:05"),
			UpdateDate:  time.Now().Format("2006-01-02 15:04:05"),
			Setor:       strconv.Itoa(saldo),
			Tarik:       "0",
			Balance:     strconv.Itoa(*data.Balance),
		}).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *EMoneyDB) InsertTransactionHistory(req *entity.TransactionHistory) error {

	if err := u.dbEMoney.Debug().Create(req).Error; err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (u *EMoneyDB) TransactionHistory(req *model.GetTrxHist) ([]*entity.TransactionHistory, error) {

	list := []*entity.TransactionHistory{}
	if err := u.dbEMoney.Debug().Where("user_id = ? and no_kartu = ?", req.UserId, req.NoKartu).
		Find(&list).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return list, nil
}
