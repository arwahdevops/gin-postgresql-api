package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       float32
	Stock       uint
	Weight      float32
	CategoryID  int            `json:"category_id"`
	Categorys   Category       `gorm:"foreignkey:CategoryID;on_delete:CASCADE"`
	Images      []ProductImage `gorm:"foreignkey:ProductID"`
}

type Category struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Name        string
	Description string
	Products    []Product `gorm:"many2many:product_categories;"`
}

type ProductImage struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	ProductID uint
	Name      string
	URL       string
}
