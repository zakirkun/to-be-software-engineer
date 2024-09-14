package models

import "time"

type Product struct {
	Id                 int     `gorm:"column:id;primaryKey"`
	ProductName        string  `gorm:"column:product_name" json:"product_name"`
	ProductImage       string  `gorm:"column:product_image" json:"product_image"`
	ProductDescription string  `gorm:"column:product_description" json:"product_description"`
	Price              float64 `gorm:"column:price" json:"price"`
	CategoryId int `gorm:"colume:category_id" json:"category_id"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (Product) TableName() string {
	return "product"
}