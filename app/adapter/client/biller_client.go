package client

import (
	"github.com/vins7/emoney-service/app/adapter/entity"
	"github.com/vins7/emoney-service/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BillerClient struct {
}

func NewBillerClient(endpoint string) *BillerClient {
	return &BillerClient{}
}

func (b *BillerClient) ListBiller() (interface{}, error) {

	client := New()
	cfg := config.GetConfig()
	res := &entity.ListBillerResponse{}
	if err := client.SetEndpoint(cfg.Client.ListBiller).
		Get(res); err != nil {
		return nil, err
	}

	if res.Code != 200 {
		return nil, status.Errorf(codes.Internal, res.Message)
	}

	return res, nil
}

func (b *BillerClient) DetailBiller(req *entity.BillerRequest) (*entity.DetailBillerResponse, error) {

	client := New()
	cfg := config.GetConfig()
	res := &entity.DetailBillerResponse{}
	if err := client.SetParams("billerId", req.ID).
		SetEndpoint(cfg.Client.DetailBiller).
		Get(res); err != nil {
		return nil, err
	}

	if res.Code != 200 {
		return nil, status.Errorf(codes.Internal, res.Message)
	}

	return res, nil
}
