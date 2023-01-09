package models

import (
	"time"

	"gorm.io/gorm"
)

// swagger:model
type CaseRecord struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Status    string
	Actors    []*Actor `gorm:"many2many:actor_in_case;"`
}

func (cr CaseRecord) GetId() string {
	return cr.ID
}
