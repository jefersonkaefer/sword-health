package grpc_notification

import (
	context "context"
	"fmt"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type NotificationClient struct {
	instance NotificationServiceClient
	addr     string
	params   keepalive.ClientParameters
}

var conn *grpc.ClientConn

func (c NotificationClient) CreateConnection(host string, port int) (*NotificationClient, error) {

	var err error

	c.addr = fmt.Sprintf("%s:%v", host, port)

	c.params = keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}

	conn, err = grpc.Dial(
		c.addr,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(
			c.params,
		),
	)

	return &c, err
}

func (c *NotificationClient) CloseConnect() {
	defer conn.Close()
}

func (c *NotificationClient) Start() NotificationServiceClient {

	c.instance = NewNotificationServiceClient(conn)

	return c.instance
}

func (c *NotificationClient) Get(userLoggedId int, id int) (notification *Notification, err error) {

	notification, err = c.instance.
		Get(
			context.Background(),
			&NotificationRequest{
				Id:           int32(id),
				UserLoggedId: int32(userLoggedId),
			},
		)

	if err != nil {
		fmt.Errorf("error: ", err)
	}
	return notification, err
}

func (c *NotificationClient) List(userLoggedId int, fromId int, limit int) (list *ListNotification, err error) {

	list, err = c.instance.
		List(
			context.Background(),
			&NotificationRequest{
				UserLoggedId: int32(userLoggedId),
				FromId:       int32(fromId),
				Limit:        int32(limit),
			},
		)

	if err != nil {
		fmt.Errorf("error: ", err)
	}

	return list, err
}
