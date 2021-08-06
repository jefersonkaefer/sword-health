package command

type TaskHandler struct {
	taskWriteService Write
	taskReadService  Read
}

func (TaskHandler) New(
	taskWriteService Write,
	taskReadService Read,
) Handler {
	return &TaskHandler{
		taskWriteService: taskWriteService,
		taskReadService:  taskReadService,
	}

}

func (uh *TaskHandler) Exec(cmd string, body []byte) {}

func (h *TaskHandler) Read() Read {

	return h.taskReadService
}

func (h *TaskHandler) Write() Write {

	return h.taskWriteService
}
