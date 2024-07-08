package postgres

import (
	"database/sql"
	pb "restaurant/genproto/users"
	"time"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) GetProfile(user *pb.UserId) (*pb.GetUser, error) {
	query :=
		`Select 
		user_name, email, phone_number 
	from 
		users 
	where 
		id = $1`

	row := u.Db.QueryRow(query, user.UserId)
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

func (U *UserRepo) UpdateProfile(profile *pb.UpdateProf) (*pb.Status, error) {
	_, err := U.Db.Exec(`UPDATE 
							users 
						SET
							user_name = $1,
							password = $2,
							email = $3,
							phone_number = $4,
							updated_at = $5
						WHERE 
							id = $6`,
		profile.Username,
		profile.Password,
		profile.Email,
		profile.Phone,
		time.Now(),
		profile.UserId)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (U *UserRepo) DeleteProfile(userId *pb.UserId) (*pb.Status, error) {
	_, err := U.Db.Exec(`UPDATE
							users
						SET
							deleted_at = $1
						WHERE
							id = $2`,
		time.Now(),
		userId.UserId)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}
