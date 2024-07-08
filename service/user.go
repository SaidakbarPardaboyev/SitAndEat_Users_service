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
	db   *sql.DB
	user *postgres.UserRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetProfile(ctx context.Context, req *pb.UserId) (*pb.GetUser, error) {
	resp, err := s.user.GetProfile(req)
	if err != nil {
		log.Fatalf("Error retrieving user information: %v", err)
		return nil, err
	}

	return resp, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, req *pb.UpdateProf) (*pb.Status, error) {
	resp, err := s.user.UpdateProfile(req)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *UserService) DeleteProfile(ctx context.Context, req *pb.UserId) (*pb.Status, error) {
	resp, err := s.user.DeleteProfile(req)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
		return nil, err
	}
	return resp, nil
}
