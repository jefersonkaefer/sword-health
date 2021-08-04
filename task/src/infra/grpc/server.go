package grpc_task

import (
	context "context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"sword-health/task/application/command"
	"sword-health/task/application/dto"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	status "google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	UnimplementedTaskServiceServer
	cmdService *command.TaskHandler
}

func (s *Server) Start(cmdService *command.TaskHandler, port int) {

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
	RegisterTaskServiceServer(gs, s)

	log.Printf("server listening at %v", lis.Addr())

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) ListTasksRequest(ctx context.Context, in *TasksListRequest) (list *TaskList, err error) {

	requestDTO := &dto.FindTask{
		OwnerId:      int(in.GetOwnerTaskId()),
		Role:         in.GetUserLoggerRole(),
		Limit:        int(in.GetLimit()),
		UserLoggedId: int(in.GetUserLoggerId()),
	}

	tasks := s.cmdService.Read().ListTasks(requestDTO)

	for _, task := range tasks {
		taskResponse := &Task{
			Id:        int32(task.ID),
			Summary:   task.Summary,
			OwnerId:   int32(task.OwnerId),
			FirstName: task.OwnerFirstName,
			LastName:  task.OwnerLastName,
			Email:     task.OwnerEmail,
			When:      task.GetWhen(),
		}
		list.Tasks = append(list.Tasks, taskResponse)
	}

	return list, nil

}

func (s *Server) FindOneTaskRequest(ctx context.Context, in *TaskRequest) (*Task, error) {

	requestDTO := &dto.FindTask{
		Id:           int(in.GetId()),
		OwnerId:      int(in.GetOwnerTaskId()),
		Role:         in.GetUserLoggerRole(),
		UserLoggedId: int(in.GetUserLoggerId()),
	}

	response := Task{}

	task, err := s.cmdService.Read().FindOne(requestDTO)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response, status.Error(http.StatusNotFound, "Task not found.")
		}
		return &response, status.Error(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	response = Task{
		Id:        int32(task.ID),
		Summary:   task.Summary,
		Status:    task.Status,
		When:      task.GetWhen(),
		FirstName: task.OwnerFirstName,
		LastName:  task.OwnerLastName,
		Email:     task.OwnerEmail,
	}

	return &response, err

}

func (s *Server) CreateTaskRequest(ctx context.Context, in *TaskRequest) (*Task, error) {

	taskRequest := dto.CreateTask(
		in.GetSummary(),
		int(in.GetOwnerTaskId()),
	)

	task, err := s.cmdService.Write().Create(taskRequest)

	if err != nil {
		return &Task{}, status.Error(http.StatusBadRequest, err.Error())
	}

	return &Task{
		Id:      int32(task.ID),
		Summary: task.Summary,
		Status:  task.Status,
		When:    task.GetWhen(),
	}, nil

}

func (s *Server) UpdateTaskRequest(ctx context.Context, in *TaskRequest) (*Task, error) {

	taskRequest := dto.TaskUpdateDTO{
		Id:             int(in.GetId()),
		Summary:        in.GetSummary(),
		Status:         in.GetStatus(),
		UserLoggedId:   int(in.GetUserLoggerId()),
		UserLoggedRole: in.GetUserLoggerRole(),
	}
	task, err := s.cmdService.Write().Update(taskRequest)

	if err != nil {
		return &Task{}, status.Error(http.StatusBadRequest, err.Error())
	}

	return &Task{
		Id:        int32(task.ID),
		Summary:   task.Summary,
		Status:    task.Status,
		When:      task.GetWhen(),
		FirstName: task.OwnerFirstName,
		LastName:  task.OwnerLastName,
		Email:     task.OwnerEmail,
	}, nil

}
