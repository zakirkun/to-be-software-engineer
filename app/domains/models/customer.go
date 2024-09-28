package models

import "time"

type Customer struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	Username  string `gorm:"unique"`
	Password  string `json:"-"`
	FullName  string
	Email     string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Customer) TableName() string {
	return "customer"
}
