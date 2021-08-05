package data_model

import "time"

type Notification struct {
	ID               uint   `gorm:"primary_key;auto_increment;not_null"`
	NotificationType string `json:"type"`
	Status           string `json:"status"`
	Content          string `json:"content"`
	FromId           int    `json:"from_id"`
	FromFullName     string `json:"from" gorm:"-"`
	When             *time.Time
}
