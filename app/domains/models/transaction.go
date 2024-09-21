package models

import (
	"time"
)

type Transaction struct {
	//Id           int        `gorm:"column:id;primaryKey" json:"-"`
	Id         int        `gorm:"column:id;primaryKey" json:"id"`
	ProductId  int        `gorm:"column:id_product" json:"id_product"`
	CustomerId int        `gorm:"column:id_customer" json:"id_customer"`
	Qty        int        `gorm:"column:qty" json:"qty"`
	Amount     float64    `gorm:"column:amount" json:"amount"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func (Transaction) TableName() string {
	return "transaction"
}

func (c Transaction) ToMapTransaction() map[string]interface{} {
	return map[string]interface{}{
		//"username":  c.Username,
		//"email":     c.Email,
		//"full_name": c.FullName,
	}

}
