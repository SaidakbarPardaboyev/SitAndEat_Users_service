package postgres

import (
	"database/sql"
	"log"
	pb "restaurant/genproto/users"
	"restaurant/models"
	"time"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) GetProfile(user *pb.UserId) (*pb.GetUser, error) {
    resp := pb.GetUser{Id: user.UserId}
    query := `
    SELECT 
        username, password, email, phone_number, 
        created_at, updated_at
    FROM 
        users 
    WHERE 
        id = $1 AND 
        deleted_at IS NULL
    `

    err := u.Db.QueryRow(query, user.UserId).Scan(&resp.Username,
        &resp.Password, &resp.Email, &resp.Phone, &resp.CreatedAt,
        &resp.UpdatedAt)

    return &resp, err
}

func (u *UserRepo) GetUserByEmail(email string) (*models.UserInfo, error) {
	query := `
	Select 
		id, username, password, phone_number 
	from 
		users 
	where 
		email = $1 and 
		deleted_at is null`

	userInfo := models.UserInfo{}
	err := u.Db.QueryRow(query, email).Scan(&userInfo.Id,
		&userInfo.Username, &userInfo.Password,
		&userInfo.Phone_number)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (U *UserRepo) StoreRefreshToken(req *models.RefreshToken) error {
	query := `
	insert into refresh_token(
		user_id, token, expires_at
	)values(
		$1, $2, $3
	)`
	_, err := U.Db.Exec(query, req.UserId, req.Token, req.ExpiresAt)
	if err != nil {
		log.Fatalf("Error with inserting refresh_token: %v", err)
		return err
	}
	return nil
}

func (U *UserRepo) UpdateProfile(profile *pb.UpdateProf) (*pb.Status, error) {
	_, err := U.Db.Exec(`UPDATE 
							users 
						SET
							username = $1,
							password = $2,
							email = $3,
							phone_number = $4,
							updated_at = $5
						WHERE 
							id = $6 and 
							deleted_at is null`,
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
							id = $2 and 
							deleted_at is null`,
		time.Now(),
		userId.UserId)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (u *UserRepo) Register(req *pb.RegisterUser) error {
	query := `
	insert into users(
		username,password,email,phone_number
	)values(
		$1,$2,$3,$4
	)`
	_, err := u.Db.Exec(query, req.Username, req.Password, req.Email, req.Phone)
	if err != nil {
		log.Fatalf("User registration error: %v", err)
		return err
	}
	return nil
}
