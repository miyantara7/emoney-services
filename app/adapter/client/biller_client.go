package client

import "github.com/vins7/emoney-service/app/adapter/entity"

type BillerClient struct {
	Endpoint string
}

func NewBillerClient(endpoint string) *BillerClient {
	return &BillerClient{
		Endpoint: endpoint,
	}
}

func (b *BillerClient) ListBiller() error {

	client := New()

	client.SetParams("billerId", "1").SetEndpoint(b.Endpoint)

	res := &entity.BillerResponse{}
	if err := client.Get(res); err != nil {
		return err
	}

	return nil
}
