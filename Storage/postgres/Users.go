package postgres

import (
	"database/sql"
	pb "restaurant/genproto/users"
	"time"
)

type NewUsers struct{
	Db *sql.DB
}

func NewUsersRepo(db *sql.DB)*NewUsers{
	return &NewUsers{Db: db}
}

func(U *NewUsers) UpdateProfile(profile *pb.UpdateProf)(*pb.Status, error){
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
	if err != nil{
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}


func(U *NewUsers) DeleteProfile(userId *pb.UserId)(*pb.Status, error){
	_, err := U.Db.Exec(`UPDATE
							users
						SET
							deleted_at = $1
						WHERE
							id = $2`, 
							time.Now(),
							userId.UserId)
	if err != nil{
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}