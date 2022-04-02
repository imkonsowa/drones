package models

import (
	"gorm.io/gorm"
	"time"
)

type AppModel struct {
	ID        uint           `json:"-" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
