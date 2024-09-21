package models

import "time"

type Transaction struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	ProductID  uint `gorm:"index;column:id_product"`
	CustomerID uint `gorm:"index;column:id_customer"`
	Qty        int
	Amount     float32
	CreatedAt  time.Time
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	Product    Product   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Customer   Customer  `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Transaction) TableName() string {
	return "transaction"
}
