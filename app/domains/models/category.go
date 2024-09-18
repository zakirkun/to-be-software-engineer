package models

import (
	"time"
)

type Category struct {
	Id           int        `gorm:"column:id;primaryKey" json:"-"`
	CategoryName string     `gorm:"column:category_name" json:"category_name"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	Products     []Product  `gorm:"foreignKey:CategoryID"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryProduct struct {
	CategoryID        uint     `gorm:"primaryKey"`
	ProductCategoryID uint     `gorm:"primaryKey"`
	Category          Category `gorm:"foreignKey:CategoryID"`
	Product           Product  `gorm:"foreignKey:ProductCategoryID"`
}
