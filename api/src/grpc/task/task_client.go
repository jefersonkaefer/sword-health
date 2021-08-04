package grpc_task

import (
	context "context"
	"fmt"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type TaskClient struct {
	instance TaskServiceClient
	addr     string
	params   keepalive.ClientParameters
}

var conn *grpc.ClientConn

func (c TaskClient) CreateConnection(host string, port int) (*TaskClient, error) {

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

func (c *TaskClient) CloseConnect() {
	defer conn.Close()
}

func (c *TaskClient) Start() TaskServiceClient {

	c.instance = NewTaskServiceClient(conn)

	return c.instance
}

func (t *TaskClient) ListTasksRequest(userId int32, role string, ownerId int32, limit int32) (*TaskList, error) {

	list, err := t.instance.ListTasksRequest(
		context.Background(),
		&TasksListRequest{
			UserLoggerId:   userId,
			UserLoggerRole: role,
			OwnerTaskId:    ownerId,
			Limit:          limit,
		},
	)

	if err != nil {
		log.Println("error: ", err)
	}

	return list, err
}

func (t *TaskClient) FindTaskRequest(
	taskId int32,
	userId int32,
	role string,
) (*Task, error) {

	task, err := t.instance.FindOneTaskRequest(
		context.Background(),
		&TaskRequest{
			Id:             taskId,
			UserLoggerId:   userId,
			UserLoggerRole: role,
		},
	)

	if err != nil {
		log.Println("error: ", err)
	}

	return task, err
}

func (t *TaskClient) CreateTaskRequest(summary string, ownerId int) (*Task, error) {

	task, err := t.instance.CreateTaskRequest(
		context.Background(),
		&TaskRequest{
			Summary:     summary,
			OwnerTaskId: int32(ownerId),
		},
	)

	if err != nil {
		log.Println("error: ", err)
	}

	return task, err
}

func (t *TaskClient) UpdateTaskRequest(taskId int, summary string, status string, userLoggerId int) (*Task, error) {

	task, err := t.instance.UpdateTaskRequest(
		context.Background(),
		&TaskRequest{
			Id:           int32(taskId),
			Summary:      summary,
			Status:       status,
			UserLoggerId: int32(userLoggerId),
		},
	)

	if err != nil {
		log.Println("error: ", err)
	}

	return task, err
}
