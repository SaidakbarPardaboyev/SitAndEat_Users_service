package main

import (
	"fmt"
	"log"
	"net"
	"restaurant/Storage/postgres"
	"restaurant/cmd/router"
	"restaurant/config"
	pb "restaurant/genproto/users"
	"restaurant/service"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().USER_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	userservice := service.NewUserService(db)
	service := grpc.NewServer()
	pb.RegisterUsersServer(service, userservice)

	go router.Router()

	fmt.Printf("Server is listening on port %s\n", config.Load().USER_SERVICE)
	if err = service.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
