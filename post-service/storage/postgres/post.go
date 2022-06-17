package postgres

import (
	pb "github.com/baxromumarov/work/post-service/genproto"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

//NewPostRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}
func (r *postRepo) GetAllData() ([]*pb.Data, error) {
	var dats []*pb.Data
	rows, err := r.db.Query("SELECT id, user_id, title, body FROM datas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var data pb.Data
		err := rows.Scan(
			&data.Id,
			&data.UserId,
			&data.Title,
			&data.Body,
		)
		if err != nil {
			return nil, err
		}
		dats = append(dats, &data)
	}

	return dats, nil
}

func (r *postRepo) GetDataById(id string) (*pb.Data, error) {
	var data pb.Data
	err := r.db.QueryRow("SELECT id, user_id, title, body FROM datas WHERE id = $1", id).Scan(
		&data.Id,
		&data.UserId,
		&data.Title,
		&data.Body,
	)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *postRepo) DeleteById(id string) error {
	_, err := r.db.Exec("DELETE FROM datas WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *postRepo) UpdateData(data *pb.Data) (*pb.Data, error) {
	_, err := r.db.Exec("UPDATE datas SET user_id = $2, title = $3, body = $4 WHERE id = $1",
		data.Id, data.UserId, data.Title, data.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
