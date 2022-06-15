package postgres

import (
	// pb "github.com/baxromumarov/work/first-service/genproto"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
    db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
    return &userRepo{db: db}
}

func (r *userRepo) CreateDB() error{
    _, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS data (
        id Integer PRIMARY KEY NOT NULL,
        user_id Integer NOT NULL,
        title text,
        body text`)
    if err != nil {
        return err
    }
    return nil
}
