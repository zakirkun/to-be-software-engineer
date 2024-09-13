package models

import (
	"time"
)

type Product struct {
	//Id           int        `gorm:"column:id;primaryKey" json:"-"`
	Id          int        `gorm:"column:id;primaryKey" json:"id"`
	CategoryId  int        `gorm:"column:category_id" json:"category_id"`
	Name        string     `gorm:"column:product_name" json:"name"`
	Image       *string    `gorm:"column:product_image" json:"image"`
	Description *string    `gorm:"column:product_description" json:"description"`
	Price       float64    `gorm:"column:price" json:"price"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (p Product) ToUpdateProduct(product Product) Product {
	updatedAt := time.Now()

	return Product{
		Id:          p.Id,
		CategoryId:  product.CategoryId,
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   &updatedAt,
	}
}

func (Product) TableName() string {
	return "product"
}
