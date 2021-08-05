package dto

type TaskCreateDTO struct {
	Summary string `json:"summary"`
	OwnerId int    `json:"owner_id"`
}

func CreateTask(
	summary string,
	ownerId int,
) TaskCreateDTO {
	return TaskCreateDTO{
		Summary: summary,
		OwnerId: ownerId,
	}
}

type TaskUpdateDTO struct {
	Id             int    `json:"id"`
	Summary        string `json:"summary"`
	Status         string `json:"status"`
	UserLoggedId   int    `json:"user_logged_id"`
	UserLoggedRole string `json:"user_logged_role"`
}

type TaskDeleteDTO struct {
	Id           int `json:"id"`
	UserLoggedId int `json:"user_logged_id"`
}

type FindTask struct {
	Id           int    `json:"id"`
	OwnerId      int    `json:"owner_id"`
	Role         string `json:"role"`
	Limit        int    `json:"limit"`
	UserLoggedId int    `json:"user_logged_id"`
}
