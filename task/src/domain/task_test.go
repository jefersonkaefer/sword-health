package domain

import (
	"sword-health/task/application/data_model"
	grpc_user "sword-health/task/infra/grpc/client/user"
	"testing"
	"time"
)

func TestCannotCreateTaskWithEmptyUser(t *testing.T) {
	_, err := Create("", 0)

	if err == nil {
		t.Error("Should not create task with empty owner.")
	}
}

func TestCannotCreateTaskWithEmptySummary(t *testing.T) {
	_, err := Create("", 1)

	if err == nil {
		t.Error("Should not create task with empty summary.")
	}
}

func TestCannotCreateTaskWithLongTextSummary(t *testing.T) {
	_, err := Create("", 2501)

	if err == nil {
		t.Error("Should not create task with longer summary.")
	}
}

func TestCannotUpdateTaskWithUserNoManagerOrNoOwner(t *testing.T) {
	user := grpc_user.User{
		Id:        1,
		FirstName: "Jonh",
		LastName:  "Doe",
		Email:     "email@mail.com",
		Role:      "tech",
		IsManager: false,
	}

	task, err := Create("task 1", 2)

	err = task.Update(int(user.GetId()), user.GetIsManager(), "summary 1")

	if err == nil {
		t.Error("Should not update task with user no manager.")
	}
}

func TestCanUpdateTaskWithOwner(t *testing.T) {
	user := grpc_user.User{
		Id:        1,
		IsManager: false,
	}

	task, err := Create("task 1", int(user.GetId()))

	err = task.Update(int(user.GetId()), user.GetIsManager(), "summary 1")

	if err != nil {
		t.Error("Should update task with owner.")
	}
}

func TestCanUpdateTaskWithManagerAndNoOwner(t *testing.T) {
	user := grpc_user.User{
		Id:        1,
		IsManager: true,
	}

	task, err := Create("task 1", 2)

	err = task.Update(int(user.GetId()), user.GetIsManager(), "summary 1")

	if err != nil {
		t.Error("Should update task with user manager.")
	}
}

func TestCanCloseTaskWithManagerAndNoOwner(t *testing.T) {
	user := grpc_user.User{
		Id:        1,
		IsManager: true,
	}

	task, err := Create("task 1", 2)

	err = task.Close(int(user.GetId()), user.GetIsManager())

	if err != nil {
		t.Error("Should close task with user manager.")
	}
}

func TestCanCloseTaskWithOwner(t *testing.T) {
	user := grpc_user.User{
		Id:        1,
		IsManager: false,
	}

	task, err := Create("task 1", int(user.GetId()))

	err = task.Close(int(user.GetId()), user.GetIsManager())

	if err != nil {
		t.Error("Should close task with owner.")
	}
}

func TestCannotDeleteTaskWithUserNoManager(t *testing.T) {
	user := grpc_user.User{
		Id:        1,
		IsManager: false,
	}

	task, err := Create("task 1", int(user.GetId()))

	err = task.Delete(user.GetIsManager())

	if err == nil {
		t.Error("Should not delete task with user no manager.")
	}
}

func TestCanDeleteTaskWithUserManager(t *testing.T) {
	user := grpc_user.User{
		Id:        1,
		IsManager: true,
	}

	task, err := Create("task 1", 2)

	err = task.Delete(user.GetIsManager())

	if err != nil {
		t.Error("Should delete task with user manager.")
	}
}

func TestIfUserIsCorrectOwnerOfTask(t *testing.T) {
	user := grpc_user.User{
		Id:        1,
		IsManager: false,
	}

	task, _ := Create("task 1", 1)

	isOwner := task.IsOwner(int(user.GetId()))

	if !isOwner {
		t.Error("This is a owner task.")
	}
}

func TestIfLoadDataCorrectFromDataModel(t *testing.T) {
	when := time.Now()
	dataModel := data_model.Task{
		ID:      1,
		OwnerId: 1,
		Summary: "suuumaty test",
		Status:  open,
		When:    &when,
	}

	task := (TaskModel{}).Load(&dataModel)

	if task.GetId() != dataModel.ID {
		t.Error("It's not the same task id.")
	}

	if int(task.GetOwnerId()) != dataModel.OwnerId {
		t.Error("It's not the same owner id.")
	}

	if task.GetSummary() != dataModel.Summary {
		t.Error("It's not the same summary.")
	}

	if task.GetWhen() != dataModel.When.Format(time.RFC822) {
		t.Error("It's not the same when date.")
	}

	if task.GetStatus() != dataModel.Status {
		t.Error("It's not the same status.")
	}

}

func TestIfLoadDataCorrectToDataModel(t *testing.T) {
	when := time.Now()

	dataModel := data_model.Task{
		ID:      1,
		OwnerId: 1,
		Summary: "suuumaty test",
		Status:  open,
		When:    &when,
	}

	task := (TaskModel{}).Load(&dataModel)

	dataModelLoad := task.GetDataModel()

	if dataModelLoad.ID != dataModel.ID {
		t.Error("It's not the same task id.")
	}

	if int(dataModelLoad.OwnerId) != dataModel.OwnerId {
		t.Error("It's not the same owner id.")
	}

	if dataModelLoad.Summary != dataModel.Summary {
		t.Error("It's not the same summary.")
	}

	if dataModelLoad.When != dataModel.When {
		t.Error("It's not the same when date.")
	}

	if dataModelLoad.Status != dataModel.Status {
		t.Error("It's not the same status.")
	}
}
