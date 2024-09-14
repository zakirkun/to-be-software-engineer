package models

import (
	"time"
)

type Customer struct {
	//Id           int        `gorm:"column:id;primaryKey" json:"-"`
	Id        int        `gorm:"column:id;primaryKey" json:"id"`
	Username  string     `gorm:"column:username" json:"username"`
	Password  string     `gorm:"column:password" json:"-"`
	FullName  string     `gorm:"column:full_name" json:"full_name"`
	Email     string     `gorm:"column:email" json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (Customer) TableName() string {
	return "customer"
}

func (c Customer) ToUpdateCustomer(name string) Customer {
	//updatedAt := time.Now()

	return Customer{
		Id:        0,
		Username:  "",
		Password:  "",
		FullName:  "",
		Email:     "",
		CreatedAt: time.Time{},
		UpdatedAt: nil,
	}

}

func (c Customer) ToMapCustomer() map[string]interface{} {
	return map[string]interface{}{
		"username":  c.Username,
		"email":     c.Email,
		"full_name": c.FullName,
	}

}
