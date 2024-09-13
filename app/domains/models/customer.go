package models

import "time"

type Customer struct {
	Id        int        `gorm:"primary_key;auto_increment" json:"-"`
	Username  string     `gorm:"username" json:"username"`
	Password  string     `gorm:"password" json:"password"`
	Email     string     `gorm:"email" json:"email"`
	FullName  string     `gorm:"full_name" json:"full_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (Customer) TableName() string {
	return "customer"
}
