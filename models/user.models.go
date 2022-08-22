package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"ID" gorm:"primarykey"`
	Name     string `json:"Name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
