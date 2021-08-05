package domain

import (
	"errors"
	"sword-health/task/application/data_model"
	"time"
)

type TaskModel struct {
	id             uint
	summary        string
	status         string
	when           *time.Time
	ownerId        int
	OwnerFirstName string
	ownerLastName  string
	email          string
}

func Create(
	summary string,
	ownerId int,
) (TaskModel, error) {

	model := TaskModel{}

	Summary, err := (Summary{}).New(summary)

	if err != nil {
		return model, err
	}

	OwnerId, err := (OwnerId{}).New(ownerId)

	if err != nil {
		return model, err
	}

	model = TaskModel{
		summary: Summary.value,
		ownerId: OwnerId.value,
		status:  open,
	}

	return model, nil
}

func (t *TaskModel) GetId() uint {
	return t.id
}

func (t *TaskModel) GetSummary() string {
	return t.summary
}

func (t *TaskModel) GetOwnerId() uint {
	return uint(t.ownerId)
}

func (t *TaskModel) GetStatus() string {
	return t.status
}

func (t *TaskModel) GetWhen() string {
	if t.when == nil {
		return ""
	}
	return t.when.Format(time.RFC822)
}

func (t *TaskModel) GetDataModel() *data_model.Task {
	data := data_model.Task{
		ID:             t.id,
		OwnerId:        t.ownerId,
		Summary:        t.summary,
		Status:         t.status,
		When:           t.when,
		OwnerFirstName: t.OwnerFirstName,
		OwnerLastName:  t.ownerLastName,
		OwnerEmail:     t.email,
	}
	return &data
}

func (TaskModel) Load(task *data_model.Task) *TaskModel {
	model := TaskModel{
		id:             task.ID,
		summary:        task.Summary,
		status:         task.Status,
		ownerId:        task.OwnerId,
		OwnerFirstName: task.OwnerFirstName,
		ownerLastName:  task.OwnerLastName,
		email:          task.OwnerEmail,
	}

	if task.When != nil {
		model.when = task.When
	}

	return &model
}

func (t *TaskModel) Update(userId int, isManager bool, summary string, status string) (err error) {

	if !t.IsOwner(userId) && !isManager {
		return errors.New("You cannot update this task.")
	}

	err = t.updateSummary(summary)

	if err != nil {
		return err
	}

	err = t.updateStatus(status)

	if err != nil {
		return err
	}

	return err
}

func (t *TaskModel) Close(userId int, isManager bool) error {

	if !t.IsOwner(userId) && !isManager {
		return errors.New("You cannot close this task.")
	}

	t.status = close

	now := time.Now()

	t.when = &now

	return nil
}

func (t *TaskModel) updateSummary(summary string) error {

	if t.summary != summary && summary != "" {
		newValue, err := (Summary{}).New(summary)
		if err != nil {
			return err
		}

		t.summary = newValue.value
	}
	return nil
}

func (t *TaskModel) updateStatus(status string) error {

	switch status {
	case open:
		t.status = status
	case close:
		now := time.Now()
		t.status = status
		t.when = &now
	default:
		return errors.New("Status invalid.")
	}

	return nil
}

func (t *TaskModel) Delete(userId int, isManager bool) error {
	if !t.IsOwner(userId) && !isManager {
		return errors.New("You cannot delete this task.")
	}
	return nil
}

func (t *TaskModel) IsOwner(userId int) bool {
	return userId == t.ownerId
}
