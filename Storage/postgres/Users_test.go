package postgres

import (
	pb "restaurant/genproto/users"
	"testing"
)


func TestUpdateProfile(t *testing.T){
	profile := &pb.UpdateProf{
		UserId: "e412e390-07ef-401f-b5c6-e931aa7ada46",
		Username: "aliyevbek",
		Password: "Aliyev#2021",
		Email: "aliyevbek@gmail.com",
		Phone: "998-90-123-4567",
	}

	db, err := ConnectDB()
	if err != nil{
		t.Error(err)
	}

	status, err := NewUsersRepo(db).UpdateProfile(profile)
	if err != nil{
		t.Error(err)
	}

	if !status.Status{
		t.Error(err)
	}
}

func TestDeleteProfile(t *testing.T){
	id := &pb.UserId{
		UserId: "e412e390-07ef-401f-b5c6-e931aa7ada46",
	}

	db, err := ConnectDB()
	if err != nil{
		t.Error(err)
	}

	status, err := NewUsersRepo(db).DeleteProfile(id)
	if err != nil{
		t.Error(err)
	}

	if !status.Status{
		t.Error(err)
	}
}