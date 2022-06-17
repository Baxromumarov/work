package main

import (
	grpcClient "github.com/baxromumarov/work/post-service/service/grpc_client"
	"google.golang.org/grpc/reflection"
	"net"

	"github.com/baxromumarov/work/post-service/config"
	pb "github.com/baxromumarov/work/post-service/genproto"
	"github.com/baxromumarov/work/post-service/pkg/db"
	"github.com/baxromumarov/work/post-service/pkg/logger"
	"github.com/baxromumarov/work/post-service/service"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "template-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	grpcC, err := grpcClient.New(cfg)
	if err != nil {
		log.Fatal("grpc client error", logger.Error(err))
		return
	}

	postService := service.NewPostService(connDB, log, grpcC)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
