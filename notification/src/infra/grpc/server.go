package grpc_notification

import (
	"fmt"
	"log"
	"net"
	"sword-health/task/application/command"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	UnimplementedNotificationServiceServer
	cmdService *command.NotificationHandler
}

func (s *Server) Start(cmdService *command.NotificationHandler, port int) {

	s.cmdService = cmdService

	var kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second,
		PermitWithoutStream: true,
	}

	var kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second,
		MaxConnectionAge:      30 * time.Second,
		MaxConnectionAgeGrace: 5 * time.Second,
		Time:                  5 * time.Second,
		Timeout:               1 * time.Second,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		fmt.Println("failed to listen: %v", err)
	}

	gs := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	RegisterNotificationServiceServer(gs, s)

	log.Printf("server listening at %v", lis.Addr())

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
