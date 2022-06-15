package repo

import (
	pb "github.com/baxromumarov/post-service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
}
