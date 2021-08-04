package command

import (
	"encoding/json"
	"log"
	"sword-health/task/application/dto"
	"sword-health/task/application/services"
)

type TaskHandler struct {
	taskWriteService *services.WriteService
	taskReadService  *services.ReadService
}

func (TaskHandler) New(
	taskWriteService *services.WriteService,
	taskReadService *services.ReadService,
) *TaskHandler {
	return &TaskHandler{
		taskWriteService: taskWriteService,
		taskReadService:  taskReadService,
	}

}

func (uh *TaskHandler) Exec(cmd string, body []byte) {
	switch cmd {
	case "task.create":
		taskDTO := dto.TaskCreateDTO{}

		if err := json.Unmarshal(body, &taskDTO); err == nil {
			uh.taskWriteService.Create(taskDTO)
		}

	default:
		log.Println("no implemented.")
	}
}

func (h *TaskHandler) Read() *services.ReadService {

	return h.taskReadService
}

func (h *TaskHandler) Write() *services.WriteService {

	return h.taskWriteService
}
