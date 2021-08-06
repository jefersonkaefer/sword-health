package data_model

import "time"

type Notification struct {
	ID               uint       `gorm:"primary_key;auto_increment;not_null"`
	NotificationType string     `json:"type"`
	Status           string     `json:"status"`
	Content          string     `json:"content"`
	FromId           int        `json:"from_id"`
	When             *time.Time `json:"when"`
	FromFullName     string     `json:"from_fullname" gorm:"-"`
}

func (n *Notification) GetWhen() string {
	if n.When == nil {
		return ""
	}
	return n.When.Format(time.RFC850)
}
