package postgres

import (
	"database/sql"
	pb "restaurant/genproto/users"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db: db}
}

func (u *Users) GetProfile(user *pb.UserId) (*pb.GetUser, error) {
	query := "Select user_name, email, phone_number from users where id = $1"

	row := u.db.QueryRow(query, user.UserId)
	var name, email, phone string
	if err := row.Scan(&name, &email, &phone); err != nil {
		return nil, err
	}
	return &pb.GetUser{
		Username: name,
		Email:    email,
		Phone:    phone,
	}, nil
}
