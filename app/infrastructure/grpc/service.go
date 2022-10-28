package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	cfg "github.com/vins7/emoney-service/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	clientBiller "github.com/vins7/emoney-service/app/adapter/client"
	db "github.com/vins7/emoney-service/app/adapter/db/e_money"
	conn "github.com/vins7/emoney-service/app/infrastructure/connection/db"
	svcBiller "github.com/vins7/emoney-service/app/service/biller"
	ucBiller "github.com/vins7/emoney-service/app/usecase/biller"
	protoBiller "github.com/vins7/module-protos/app/interface/grpc/proto/e_money_service"
)

func RunServer() {

	config := cfg.GetConfig()
	grpcServer := grpc.NewServer()

	Apply(grpcServer)
	reflection.Register(grpcServer)

	svcHost := config.Server.Grpc.Host
	svcPort := config.Server.Grpc.Port

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", svcHost, svcPort))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start notpool Service gRPC server: %v", err)
		}
	}()

	fmt.Printf("gRPC server is running at %s:%d\n", svcHost, svcPort)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())

}

func Apply(server *grpc.Server) {
	protoBiller.RegisterBillerServiceServer(server,
		svcBiller.NewBillerService(
			*ucBiller.NewBillerUsecase(
				clientBiller.NewBillerClient("https://phoenix-imkas.ottodigital.id/interview/biller/v1/detail"),
				db.NewEMoneyDB(conn.UserDB, conn.EMoneyDB),
			),
		),
	)
}
