package grpc_user

import (
	context "context"
	"fmt"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type UserClient struct {
	instance UserServiceClient
	addr     string
	params   keepalive.ClientParameters
}

var conn *grpc.ClientConn

func (c UserClient) CreateConnection(host string, port int) (*UserClient, error) {

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

func (c *UserClient) CloseConnect() {
	defer conn.Close()
}

func (c *UserClient) Start() UserServiceClient {

	c.instance = NewUserServiceClient(conn)

	return c.instance
}

func (c *UserClient) Get(userId int) (user *User, err error) {

	user, err = c.instance.
		Get(
			context.Background(),
			&User{Id: int32(userId)},
		)

	if err != nil {
		fmt.Errorf("error: ", err)
	}
	return user, err
}
