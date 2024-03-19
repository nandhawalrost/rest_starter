package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string
	Quantity    uint
	Active      bool `gorm:"default:false"`
}
