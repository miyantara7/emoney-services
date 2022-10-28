package client

import "github.com/vins7/emoney-service/app/adapter/entity"

type BillerRepo interface {
	ListBiller() (interface{}, error)
	DetailBiller(req *entity.BillerRequest) (*entity.DetailBillerResponse, error)
}
