package service

import (
	"context"

	pb "github.com/baxromumarov/work/first-service/genproto"
	l "github.com/baxromumarov/work/first-service/pkg/logger"
	"github.com/baxromumarov/work/first-service/storage"
	"github.com/jmoiron/sqlx"
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

// func (s *UserService) CreateDB(ctx context.Context, *pb.Empty) error {
    
// }
