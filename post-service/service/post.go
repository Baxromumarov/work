package service

import (
	"context"
	pb "github.com/baxromumarov/work/post-service/genproto"
	l "github.com/baxromumarov/work/post-service/pkg/logger"
	cl "github.com/baxromumarov/work/post-service/service/grpc_client"
	"github.com/baxromumarov/work/post-service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *PostService) GetAllData(ctx context.Context, req *pb.Empty) (*pb.DataResp, error) {
	datas, err := s.storage.Post().GetAllData()
	if err != nil {
		s.logger.Error("Error while getting data list", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting data list")
	}
	return &pb.DataResp{
		Data: datas,
	}, nil
}

func (s *PostService) GetDataById(ctx context.Context, req *pb.ByIdReq) (*pb.Data, error) {
	data, err := s.storage.Post().GetDataById(req.Id)
	if err != nil {
		s.logger.Error("Error while getting databy id ", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting data by id")
	}
	return data, nil
}

func (s *PostService) DeleteById(ctx context.Context, req *pb.ByIdReq) (*pb.Empty, error) {
	err := s.storage.Post().DeleteById(req.Id)
	if err != nil {
		s.logger.Error("Error while deleting data", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while deleting data")
	}
	return &pb.Empty{}, nil
}

func (s *PostService) UpdateData(ctx context.Context, req *pb.Data) (*pb.Data, error) {
	data, err := s.storage.Post().UpdateData(req)
	if err != nil {
		s.logger.Error("Error while updating data", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while updating data")
	}
	return data, nil
}
