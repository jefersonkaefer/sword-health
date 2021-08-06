package domain

import (
	"errors"
	"sword-health/notification/application/data_model"
	"time"
)

const (
	unread = "unread"
	read   = "read"
)

type NotificationModel struct {
	id               uint
	notificationType string
	status           string
	content          string
	fromId           int
	FromFullName     string
	when             *time.Time
}

func Create(
	notificationType string,
	content string,
	fromId int,
) (NotificationModel, error) {
	now := time.Now()
	return NotificationModel{
		notificationType: notificationType,
		content:          content,
		fromId:           fromId,
		status:           unread,
		when:             &now,
	}, nil

}

func (n *NotificationModel) GetId() uint {
	return n.id
}

func (n *NotificationModel) GetFromId() int {
	return n.fromId
}

func (n *NotificationModel) GetStatus() string {
	return n.status
}

func (n *NotificationModel) GetWhen() string {
	if n.when == nil {
		return ""
	}
	return n.when.Format(time.RFC822)
}

func (n *NotificationModel) GetDataModel() *data_model.Notification {
	data := data_model.Notification{
		ID:      n.id,
		Status:  n.status,
		When:    n.when,
		Content: n.content,
		FromId:  n.fromId,
	}
	return &data
}
func (NotificationModel) Load(notification *data_model.Notification) *NotificationModel {
	model := NotificationModel{
		id:               notification.ID,
		notificationType: notification.NotificationType,
		content:          notification.Content,
		status:           notification.Status,
		fromId:           notification.FromId,
	}

	if notification.When != nil {
		model.when = notification.When
	}

	return &model
}

func (n *NotificationModel) MarkAsRead(IsManager bool) (err error) {

	if !IsManager {
		return errors.New("You cannot access this notification.")
	}

	n.status = read

	return err
}
