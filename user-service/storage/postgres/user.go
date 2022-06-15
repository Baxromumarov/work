package postgres

import (
	// "fmt"

	pb "github.com/baxromumarov/work/user-service/genproto"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}



func (r *userRepo) Create(req *pb.Request ) (*pb.Empty, error) {
	
	for _, val := range req.Data {
		_,err := r.db.Exec(`INSERT INTO datas (id, user_id, title, body ) VALUES ($1, $2, $3, $4)`,
			val.Id, val.UserId, val.Title, val.Body) 
		if err != nil {
			return nil, err
				
		}
	}
	
	_, err := r.db.Exec(`INSERT INTO paginations (total, pages, page ) VALUES ($1, $2, $3)`,
		req.Meta.Pagination.Total, 
		req.Meta.Pagination.Pages, 
		req.Meta.Pagination.Page, 
		
	)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(`INSERT INTO links (previous, current, next ) VALUES ($1, $2, $3)`,
		req.Meta.Pagination.Links.Previous, 
		req.Meta.Pagination.Links.Current, 
		req.Meta.Pagination.Links.Next,
	)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
