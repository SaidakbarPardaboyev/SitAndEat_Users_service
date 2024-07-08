package postgres

import (
	"database/sql"
	pb "restaurant/genproto/users"
	"testing"
)

func TestGetProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	User := pb.GetUser{
		Username: "aliyevbek",
		Email:    "aliyevbek@gmail.com",
		Phone:    "998-90-123-4567",
	}

	users := NewUserRepo(db)
	res, err := users.GetProfile(&pb.UserId{UserId: "8e52ea77-abe1-49c2-8160-8eaa9b2d89b3"})

	if err != nil || err == sql.ErrNoRows {
		t.Fatalf("Failed to get profile: %v", err)
	}
	if res.Username != User.Username {
		t.Fatalf("Failed to get profile: %v", err)
	}
	if res.Email != User.Email {
		t.Fatalf("Failed to get profile: %v", err)
	}
	if res.Phone != User.Phone {
		t.Fatalf("Failed to get profile: %v", err)
	}
}

func TestUpdateProfile(t *testing.T) {
	profile := &pb.UpdateProf{
		UserId:   "e412e390-07ef-401f-b5c6-e931aa7ada46",
		Username: "aliyevbek",
		Password: "Aliyev#2021",
		Email:    "aliyevbek@gmail.com",
		Phone:    "998-90-123-4567",
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}

	status, err := NewUserRepo(db).UpdateProfile(profile)
	if err != nil {
		t.Error(err)
	}

	if !status.Status {
		t.Error(err)
	}
}

func TestDeleteProfile(t *testing.T) {
	id := &pb.UserId{
		UserId: "e412e390-07ef-401f-b5c6-e931aa7ada46",
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}

	status, err := NewUserRepo(db).DeleteProfile(id)
	if err != nil {
		t.Error(err)
	}

	if !status.Status {
		t.Error(err)
	}
}
