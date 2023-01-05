package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Name        string
	Description string
	Price       float32
	Stock       uint
	Category    Category `gorm:"foreignkey:CategoryID"`
	CategoryID  uint
	Images      []ProductImage `gorm:"foreignkey:ProductID"`
}

type Category struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"`
}

type ProductImage struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	ProductID uint
	URL       string
}
