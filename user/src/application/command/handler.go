package command

import (
	"encoding/json"
	"log"
	"sword-health/user/application/dto"
	"sword-health/user/application/services"
)

type UserHandler struct {
	userWriteService *services.WriteService
	userReadService  *services.ReadService
}

func (UserHandler) New(
	userWriteService *services.WriteService,
	userReadService *services.ReadService,
) *UserHandler {
	return &UserHandler{
		userWriteService: userWriteService,
		userReadService:  userReadService,
	}

}

func (uh *UserHandler) Exec(cmd string, body []byte) {
	switch cmd {
	case "user.create":
		userDTO := dto.UserCreateDTO{}

		if err := json.Unmarshal(body, &userDTO); err == nil {
			uh.userWriteService.Create(userDTO)
		}

	default:
		log.Println("no implemented.")
	}
}

func (h *UserHandler) Read() *services.ReadService {

	return h.userReadService
}

func (h *UserHandler) Write() *services.WriteService {

	return h.userWriteService
}
