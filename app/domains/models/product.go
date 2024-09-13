package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	Id                 int             `gorm:"primary_key;auto_increment" json:"-"`
	CategoryId         int             `gorm:"category_id;unique" json:"category_id"`
	ProductName        string          `gorm:"product_name" json:"product_name"`
	ProductImage       string          `gorm:"product_image" json:"product_image"`
	ProductDescription string          `gorm:"product_description" json:"product_description"`
	Price              decimal.Decimal `gorm:"price" json:"price"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          *time.Time      `json:"updated_at"`
}

func (Product) TableName() string {
	return "product"
}
