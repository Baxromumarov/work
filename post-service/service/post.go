package service

import (
	"context"
	pb "github.com/baxromumarov/work/post-service/genproto"
	l "github.com/baxromumarov/work/post-service/pkg/logger"
	cl "github.com/baxromumarov/work/post-service/service/grpc_client"
	"github.com/baxromumarov/work/post-service/storage"
	"github.com/jmoiron/sqlx"
)

//PostService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  cl.GrpcClientI
}

//NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client cl.GrpcClientI) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}
func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	return nil, nil
}
