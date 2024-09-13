package models

import (
	"time"
)

type Category struct {
	//Id           int        `gorm:"column:id;primaryKey" json:"-"`
	Id           int        `gorm:"column:id;primaryKey" json:"id"`
	CategoryName string     `gorm:"column:category_name" json:"category_name"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "category"
}

func (c Category) ToUpdateCategory(name string) Category {
	updatedAt := time.Now()

	return Category{
		Id:           c.Id,
		CategoryName: name,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    &updatedAt,
	}

}
