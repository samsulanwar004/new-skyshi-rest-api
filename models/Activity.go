package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        int             `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title     string          `gorm:"column:title" json:"title"`
	Email     string          `gorm:"column:email" json:"email"`
	CreatedAt time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}
type ActivityNull struct {
}
