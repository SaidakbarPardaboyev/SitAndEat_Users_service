package postgres

import (
	"database/sql"
	"log"
	pb "restaurant/genproto/users"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) RegisterUser(req *pb.RegisterUser) (*pb.Status, error) {
	query := `
	insert into users(
		username,password,email,phone_number
	)values(
		$1,$2,$3,$4
	)`
	_, err := u.db.Exec(query, req.Username, req.Password, req.Email, req.Phone)
	if err != nil {
		log.Fatalf("User registration error: %v", err)
		return nil, err
	}
	return &pb.Status{Status: true}, nil
}
