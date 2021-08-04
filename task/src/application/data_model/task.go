package data_model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID             uint           `gorm:"primary_key;auto_increment;not_null"`
	OwnerId        int            `json:"owner_id"`
	Summary        string         `json:"summary"`
	Status         string         `json:"status"`
	When           *time.Time     `json:"when,omitempty"`
	OwnerFirstName string         `json:"first_name" gorm:"-"`
	OwnerLastName  string         `json:"last_name" gorm:"-"`
	OwnerEmail     string         `json:"ow_email" gorm:"-"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at omitempty" gorm:"index"`
}

func (t *Task) GetWhen() string {
	if t.When == nil {
		return ""
	}
	return t.When.Format(time.RFC822)
}
