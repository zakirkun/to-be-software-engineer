package models

import "time"

type Customer struct {
	Id        int     `gorm:"column:id;primaryKey"`
	Username  string  `gorm:"column:username;unique" json:"username"`
	Password  string  `gorm:"column:password" json:"password"`
	FullName  string  `gorm:"column:full_name" json:"full_name"`
	Email     string `gorm:"column:email;unique" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (Customer) TableName() string {
	return "customer"
}