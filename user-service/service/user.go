package service

import (
	"context"

	pb "github.com/baxromumarov/work/user-service/genproto"
	l "github.com/baxromumarov/work/user-service/pkg/logger"
	"github.com/baxromumarov/work/user-service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
    storage storage.IStorage
    logger  l.Logger
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
    return &UserService{
        storage: storage.NewStoragePg(db),
        logger:  log,
    }
}


func (s *UserService) Create(ctx context.Context, req *pb.Request) (*pb.Empty, error){
    info, err := s.storage.User().Create(req)
    if err != nil {
        s.logger.Error("Error while inserting informations to db", l.Error(err))
        return nil, status.Error(codes.Internal,"Error while inserting informations to db")
    }
    

    return info,nil
}
