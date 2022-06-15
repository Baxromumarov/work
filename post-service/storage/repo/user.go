package repo

import (
    pb "github.com/rustagram/template-service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
    Create(*pb.User) (*pb.User, error)
}
