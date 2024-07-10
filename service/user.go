package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"restaurant/Storage/postgres"
	pb "restaurant/genproto/users"
	"restaurant/pkg/logger"
)

type UserService struct {
	pb.UnimplementedUsersServer
	Logger *slog.Logger
	User   *postgres.UserRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		User:   postgres.NewUserRepo(db),
		Logger: logger.NewLogger(),
	}
}

func (U *UserService) GetProfile(ctx context.Context, req *pb.UserId) (*pb.GetUser, error) {
	resp, err := U.User.GetProfile(req)
	if err != nil {
		U.Logger.Error(fmt.Sprintf("Error retrieving user information: %v", err))
		return nil, err
	}

	return resp, nil
}

func (U *UserService) UpdateProfile(ctx context.Context, req *pb.UpdateProf) (*pb.Status, error) {
	resp, err := U.User.UpdateProfile(req)
	if err != nil {
		U.Logger.Error(fmt.Sprintf("Error updating user: %v", err))
		return nil, err
	}
	return resp, nil
}

func (U *UserService) DeleteProfile(ctx context.Context, req *pb.UserId) (*pb.Status, error) {
	resp, err := U.User.DeleteProfile(req)
	if err != nil {
		U.Logger.Error(fmt.Sprintf("Error deleting user: %v", err))
		return nil, err
	}
	return resp, nil
}
