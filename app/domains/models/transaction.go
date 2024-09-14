package models

import "github.com/shopspring/decimal"

type Transaction struct {
	Id         int             `gorm:"column:id;primaryKey" json:"-"`
	IdProduct  int             `gorm:"column:id_product" json:"-"`
	IdCustomer int             `gorm:"column:id_customer" json:"-"`
	Qty        int             `gorm:"column:qty" json:"qty"`
	Amount     decimal.Decimal `gorm:"column:amount" json:"amount"`
}

func (Transaction) TableName() string {
	return "transaction"
}
