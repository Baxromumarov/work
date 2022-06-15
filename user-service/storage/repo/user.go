package repo

import(
	pb "github.com/baxromumarov/work/user-service/genproto"
)
//UserStorageI ...
type UserStorageI interface {
	Create(*pb.Request) (*pb.Empty, error)
}
