package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	grpc_task "sword-health/api/grpc/task"
	"sword-health/api/infra/amqp"

	middleware "sword-health/api/http"
	"sword-health/api/validators"

	"github.com/gin-gonic/gin"
	status "google.golang.org/grpc/status"
)

type TaskController struct {
	Validator  *validators.JSONValidator
	TaskClient *grpc_task.TaskClient
	AMQ        *amqp.Connection
}

type TaskCreateRequest struct {
	Summary string `json:"summary" binding:"required"`
}

type TaskUpdateRequest struct {
	Summary string `json:"summary" binding:"required" validate:"max=2500"`
	Status  string `json:"status" binding:"required" validate:"alpha"`
}

func (t *TaskController) Create(c *gin.Context) {

	userLogged, _ := c.Get("userLogged")
	user := userLogged.(*middleware.UserLogged)

	var request TaskCreateRequest

	if errors := t.Validator.Validate(c, &request); errors != nil {
		return
	}

	task, err := t.TaskClient.CreateTaskRequest(request.Summary, int(user.Id))

	if err != nil {
		status, _ := status.FromError(err)
		c.JSON(int(status.Code()), gin.H{"error": status.Message()})
		return
	}
	// message := struct {
	// 	OwnerId int    `json:"owner_id"`
	// 	Summary string `json:"summary"`
	// }{
	// 	OwnerId: int(user.Id),
	// 	Summary: request.Summary,
	// }

	// if message, err := json.Marshal(message); err == nil {
	// 	t.AMQ.Dispatch(amqp.ExchangeTask, amqp.RouteKeyTaskCreate, message)
	// }

	c.JSON(http.StatusCreated, task)
}

func (t *TaskController) Update(c *gin.Context) {

	userLogged, _ := c.Get("userLogged")
	user := userLogged.(*middleware.UserLogged)

	taskId := c.Param("id")
	id, _ := strconv.Atoi(taskId)

	var request TaskUpdateRequest

	if errors := t.Validator.Validate(c, &request); errors != nil {
		return
	}

	task, err := t.TaskClient.UpdateTaskRequest(
		id,
		request.Summary,
		request.Status,
		int(user.Id),
	)

	if err != nil {
		status, _ := status.FromError(err)
		c.JSON(int(status.Code()), gin.H{"error": status.Message()})
		return
	}
	// message := struct {
	// 	OwnerId int    `json:"owner_id"`
	// 	Summary string `json:"summary"`
	// }{
	// 	OwnerId: int(user.Id),
	// 	Summary: request.Summary,
	// }

	// if message, err := json.Marshal(message); err == nil {
	// 	t.AMQ.Dispatch(amqp.ExchangeTask, amqp.RouteKeyTaskCreate, message)
	// }

	c.JSON(http.StatusCreated, task)
}

func (t *TaskController) Delete(c *gin.Context) {
	// id := c.Param("id")
}

func (t *TaskController) List(c *gin.Context) {

	userLogged, _ := c.Get("userLogged")
	user := userLogged.(*middleware.UserLogged)

	owner := c.Query("owner_id")
	ownerId, _ := strconv.Atoi(owner)

	limitParam := c.Query("limit")
	limit, _ := strconv.Atoi(limitParam)

	list, err := t.TaskClient.ListTasksRequest(
		user.Id,
		user.Role,
		int32(ownerId),
		int32(limit),
	)

	if err != nil {
		status, _ := status.FromError(err)
		c.JSON(int(status.Code()), gin.H{"error": status.Message()})
		fmt.Println(err)
		return
	}

	if list.Tasks == nil {
		c.JSON(http.StatusNoContent, http.StatusText(http.StatusNoContent))
		return
	}

	c.JSON(http.StatusOK, list)
}

func (t *TaskController) Get(c *gin.Context) {

	userLogged, _ := c.Get("userLogged")
	user := userLogged.(*middleware.UserLogged)

	taskId := c.Param("id")
	id, _ := strconv.Atoi(taskId)

	task, err := t.TaskClient.FindTaskRequest(
		int32(id),
		user.Id,
		user.Role,
	)

	if err != nil {
		status, _ := status.FromError(err)
		c.JSON(int(status.Code()), gin.H{"error": status.Message()})
		return
	}

	c.JSON(http.StatusOK, task)
}
