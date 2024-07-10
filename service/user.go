package service

import (
	"context"
	"database/sql"
	"log"
	"restaurant/Storage/postgres"
	pb "restaurant/genproto/users"
)

type UserService struct {
	pb.UnimplementedUsersServer
	User *postgres.UserRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{User: postgres.NewUserRepo(db)}
}

func (U *UserService) GetProfile(ctx context.Context, req *pb.UserId) (*pb.GetUser, error) {
	resp, err := U.User.GetProfile(req)
	if err != nil {
		log.Fatalf("Error retrieving user information: %v", err)
		return nil, err
	}

	return resp, nil
}

func (U *UserService) UpdateProfile(ctx context.Context, req *pb.UpdateProf) (*pb.Status, error) {
	resp, err := U.User.UpdateProfile(req)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
		return nil, err
	}
	return resp, nil
}

func (U *UserService) DeleteProfile(ctx context.Context, req *pb.UserId) (*pb.Status, error) {
	resp, err := U.User.DeleteProfile(req)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
		return nil, err
	}
	return resp, nil
}
