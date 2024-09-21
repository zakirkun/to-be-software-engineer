package models

import (
	"time"
)

type Product struct {
	Id                 int        `gorm:"column:id;primaryKey" json:"-"`
	CategoryId         int        `gorm:"column:category_id" json:"category_id"`
	ProductName        string     `gorm:"column:product_name" json: "product_name"`
	ProductImage       string     `gorm:"column:product_image" json: "product_image"`
	ProductDescription string     `gorm:"column:product_description" json: "product_description"`
	Price              float32    `gorm:"column:price" json: "price"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
}

func (Product) TableName() string {
	return "product"
}
