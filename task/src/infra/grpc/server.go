package grpc_task

import (
	context "context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"sword-health/task/application/command"
	"sword-health/task/application/data_model"
	"sword-health/task/application/dto"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	status "google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	UnimplementedTaskServiceServer
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

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
	list = &TaskList{}

	listTasks, err := s.cmdService.Read().ListTasks(
		int(in.GetUserLoggedId()),
		int(in.GetOwnerTaskId()),
		int(in.GetLimit()),
	)

	for _, task := range listTasks {
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

func (s *Server) FindOneTaskRequest(ctx context.Context, in *TaskRequest) (task *Task, err error) {

	var taskDataModel *data_model.Task

	taskDataModel, err = s.cmdService.Read().FindOne(
		int(in.GetUserLoggedId()),
		int(in.GetId()),
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return task, status.Error(http.StatusNotFound, "Task not found.")
		}
		return task, status.Error(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return &Task{
		Id:        int32(taskDataModel.ID),
		Summary:   taskDataModel.Summary,
		Status:    taskDataModel.Status,
		When:      taskDataModel.GetWhen(),
		FirstName: taskDataModel.OwnerFirstName,
		LastName:  taskDataModel.OwnerLastName,
		Email:     taskDataModel.OwnerEmail,
	}, err

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
		UserLoggedId:   int(in.GetUserLoggedId()),
		UserLoggedRole: in.GetUserLoggedRole(),
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

func (s *Server) DeleteTaskRequest(ctx context.Context, in *TaskRequest) (*Task, error) {

	err := s.cmdService.Write().Delete(
		int(in.GetUserLoggedId()),
		int(in.GetId()),
	)

	if err != nil {
		return &Task{}, status.Error(http.StatusBadRequest, err.Error())
	}

	return &Task{}, nil

}

func (s *Server) CloseTaskRequest(ctx context.Context, in *TaskRequest) (task *Task, err error) {

	err = s.cmdService.Write().Close(
		int(in.GetUserLoggedId()),
		int(in.GetId()),
	)

	if err != nil {
		return &Task{}, status.Error(http.StatusBadRequest, err.Error())
	}

	return &Task{}, nil

}
