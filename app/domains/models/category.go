package models

import (
	"time"
)

type Category struct {
	Id           int        `gorm:"column:id;primaryKey" json:"id"`
	CategoryName string     `gorm:"column:category_name" json:"category_name"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "category"
}
