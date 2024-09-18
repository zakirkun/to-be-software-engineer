package models

import "time"

type Product struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	CategoryID         uint      `gorm:"index" json:"-"`
	ProductName        string    `json:"product_name"`
	ProductImage       string    `json:"product_image"`
	ProductDescription string    `json:"product_description"`
	Price              float32   `json:"price"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Category           Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Product) TableName() string {
	return "product"
}
