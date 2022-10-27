package main

import (
	grpc "github.com/vins7/emoney-service/app/infrastructure/grpc"
)

func main() {
	grpc.RunServer()
}
