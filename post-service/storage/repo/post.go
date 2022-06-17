package repo

import pb "github.com/baxromumarov/work/post-service/genproto"

//PostStorageI ...
type PostStorageI interface {
	GetAllData() ([]*pb.Data, error)
	GetDataById(id string) (*pb.Data, error)
	DeleteById(id string) error
	UpdateData(*pb.Data) (*pb.Data, error)
}
