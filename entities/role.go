package entities

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID             uint            `gorm:"primary_key" json:"id"`
	Name           string          `gorm:"size:255;not null" json:"name"`
	SlugName       string          `gorm:"size:255;not null;index" json:"guard_name"`
	Description    string          `gorm:"size:255;" json:"description"`
	Authorizations []Authorization `gorm:"many2many:role_authorizations;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"authorizations"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      gorm.DeletedAt
}

type Roles []Role

func (r Roles) Origin() []Role {
	return []Role(r)
}
