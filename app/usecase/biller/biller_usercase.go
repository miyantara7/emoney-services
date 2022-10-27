package biller

import (
	"github.com/mitchellh/mapstructure"
	repo "github.com/vins7/emoney-service/app/adapter/client"
	"github.com/vins7/emoney-service/app/adapter/entity"
)

type BillerUsecase struct {
	repo repo.BillerRepo
}

func NewBillerUsecase(repo repo.BillerRepo) *BillerUsecase {
	return &BillerUsecase{
		repo: repo,
	}
}

func (u *BillerUsecase) ListBiller(in interface{}) (interface{}, error) {
	var req *entity.User

	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, err
	}

	if err := u.repo.ListBiller(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *BillerUsecase) DetailBiller(in interface{}) (interface{}, error) {
	var req *entity.User

	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, err
	}

	if err := u.repo.ListBiller(); err != nil {
		return nil, err
	}

	return nil, nil
}
