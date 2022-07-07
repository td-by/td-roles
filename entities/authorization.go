package entities

import (
	"gorm.io/gorm"
	"time"
)

type Authorization struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	SlugName    string    `gorm:"size:255;not null;index" json:"guard_name"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt
}

type Authorizations []Authorization

func (a Authorizations) Origin() []Authorization {
	return []Authorization(a)
}
