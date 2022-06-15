package repo

import (
	// pb "github.com/baxromumarov/work/first-service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	CreateDB() error
}
