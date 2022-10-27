package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BuilderResty struct {
	Endpoint    string
	RestyClient *resty.Client
}

func New() *BuilderResty {
	return &BuilderResty{
		RestyClient: resty.New(),
	}
}

func (r *BuilderResty) SetParams(key, val string) *BuilderResty {
	r.RestyClient = r.RestyClient.SetQueryParam(key, val)
	return r
}

func (r *BuilderResty) SetEndpoint(endpoint string) *BuilderResty {
	r.Endpoint = endpoint
	return r
}

func (b *BuilderResty) Get(response interface{}) error {
	fmt.Println("Endpoint : ", b.Endpoint)
	data, err := b.RestyClient.R().Get(b.Endpoint)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	fmt.Println("DATA : ", data)
	return nil
}
