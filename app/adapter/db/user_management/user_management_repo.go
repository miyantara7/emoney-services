package user_management

import (
	"github.com/vins7/emoney-service/app/adapter/entity"
	"github.com/vins7/emoney-service/app/interface/model"
)

type UserManagementRepo interface {
	Login(interface{}) (*entity.User, error)
	CreateUser(interface{}) error
	InsertLoginHistory(interface{}) error
	DetailUser(*model.DetailUserReq) (*entity.User, error)
}
