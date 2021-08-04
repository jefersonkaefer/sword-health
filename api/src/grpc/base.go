package base_grpc

import (
	"fmt"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Client struct {
	addr   string
	params keepalive.ClientParameters
}

func (c Client) CreateConnection(host string, port int) (conn *grpc.ClientConn, err error) {

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

	return conn, err
}
