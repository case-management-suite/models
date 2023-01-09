package models

import (
	"time"

	"gorm.io/gorm"
)

// swagger:model
type CaseAction struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CaseRecord   CaseRecord
	CaseRecordID string `json:"case_record_id"`
	Action       string
}
