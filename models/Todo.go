package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID              int             `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ActivityGroupID int             `gorm:"column:activity_group_id;index" json:"activity_group_id"`
	Title           string          `gorm:"column:title" json:"title"`
	IsActive        bool            `gorm:"column:is_active" json:"is_active"`
	Priority        string          `gorm:"column:priority" json:"priority"`
	CreatedAt       time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       *gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}
type TodoNull struct {
}
