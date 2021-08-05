package grpc_user

import (
	context "context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"sword-health/user/application/command"
	"sword-health/user/application/dto"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	status "google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	UnimplementedUserServiceServer
	command *command.UserHandler
}

func (s *Server) Start(command *command.UserHandler, port int) {

	s.command = command

	var kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	var kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		fmt.Println("failed to listen: %v", err)
	}

	gs := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	RegisterUserServiceServer(gs, s)

	log.Printf("server listening at %v", lis.Addr())

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) CheckUser(ctx context.Context, in *CheckUserRequest) (user *User, err error) {

	userModel := s.command.Read().FindByEmail(in.Email)

	if !userModel.CheckEmail(in.Email) || !userModel.CheckPassword(in.Password) {
		return user, status.Error(http.StatusBadRequest, "Invalid email or password.")
	}

	user = &User{
		Id:        int32(userModel.GetId()),
		FirstName: userModel.GetFirstName(),
		LastName:  userModel.GetLastName(),
		Email:     userModel.GetEmail(),
		Role:      userModel.GetRole(),
	}

	return user, nil

}

func (s *Server) Get(ctx context.Context, in *User) (user *User, err error) {

	userModel, err := s.command.Read().FindOne(int(in.GetId()))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, status.Error(http.StatusNotFound, "Task not found.")
		}
		return user, status.Error(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	user = &User{
		Id:        int32(userModel.GetId()),
		FirstName: userModel.GetFirstName(),
		LastName:  userModel.GetLastName(),
		Email:     userModel.GetEmail(),
		Role:      userModel.GetRole(),
	}

	return user, nil

}

func (s *Server) CreateUser(ctx context.Context, in *CreateUserRequest) (*User, error) {

	userDTO := dto.CreateUser(
		in.GetFirstName(),
		in.GetLastName(),
		in.GetEmail(),
		in.GetRole(),
		in.GetPassword(),
		in.GetRePassword(),
	)

	user, err := s.command.Write().Create(userDTO)

	if err != nil {
		return &User{}, status.Error(http.StatusBadRequest, err.Error())
	}

	return &User{
		Id: int32(user.ID),
	}, err
}
