package data_model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary_key;auto_increment;not_null"`
	FirstName string `gorm:"size:60"`
	LastName  string `gorm:"size:60"`
	Email     string `gorm:"uniqueIndex;size:255"`
	Password  string `gorm:";size:255"`
	Role      string
	IsDeleted bool
	CreatedBy int
	UpdatedBy int
	DeletedBy int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
