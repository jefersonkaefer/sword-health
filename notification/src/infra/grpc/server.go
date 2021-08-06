package grpc_notification

import (
	context "context"
	"fmt"
	"log"
	"net"
	"sword-health/notification/application/command"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	UnimplementedNotificationServiceServer
	cmdService command.Handler
}

func (s *Server) Start(cmdService command.Handler, port string) {

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

func (s *Server) Get(ctx context.Context, in *NotificationRequest) (notification *Notification, err error) {
	
	notificationResponse := &Notification{}

	dataModel, err := s.cmdService.Read().FindOne(
		int(in.GetUserLoggedId()),
		int(in.GetId()),
	)

	if err != nil{
		return notificationResponse, err
	}
	
	go s.cmdService.Write().MarkAsRead(
		int(in.GetUserLoggedId()),
		int(dataModel.ID),
	)

	notificationResponse = &Notification{
		Id:               int32(dataModel.ID),
		NotificationType: dataModel.NotificationType,
		Content:          dataModel.Content,
		FromFullName:     dataModel.FromFullName,
		When:             dataModel.GetWhen(),
	}

	return notificationResponse, err
}

func (s *Server) List(ctx context.Context, in *NotificationRequest) (list *ListNotification, err error) {
	list = &ListNotification{}

	notifications, err := s.cmdService.Read().ListNotifications(
		int(in.GetUserLoggedId()),
		int(in.GetFromId()),
		int(in.GetLimit()),
	)
	if err != nil {
		return list, err
	}

	for _, notification := range notifications {
		list.Notifications = append(list.Notifications, &Notification{
			Id:               int32(notification.ID),
			NotificationType: notification.NotificationType,
			Content:          notification.Content,
			FromFullName:     notification.FromFullName,
			Status:           notification.Status,
			When:             notification.GetWhen(),
		})
	}

	return list, err
}
