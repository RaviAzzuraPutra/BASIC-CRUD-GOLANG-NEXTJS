package model

import (
	"time"

	"gorm.io/gorm"
)

type Travel struct {
	Id          *string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Photo       *string `json:"photo"`
	Price       *int    `json:"price"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
