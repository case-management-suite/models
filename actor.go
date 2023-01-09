package models

import (
	"time"

	"gorm.io/gorm"
)

//go:generate go run ./main/gen.go --types=Actor,CaseContext,CaseRecord,CaseAction
type Actor struct {
	ID            string `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	CasesInvolved []*CaseRecord  `gorm:"many2many:actor_in_case;"`
}
