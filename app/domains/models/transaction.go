package models

import (
	"time"
)

type Transaction struct {
	Id           int        `gorm:"column:id;primaryKey" json:"-"`
	ProductId int `gorm:"column:id_product" json:"product_id"`
	CustomerId int `gorm:"column:id_customer" json:"customer_id"`
	Qty int `gorm:"column:qty" json:"qty"`
	Amount float64 `gorm:"column:amount" json:"amount"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

func (Transaction) TableName() string {
	return "transaction"
}
