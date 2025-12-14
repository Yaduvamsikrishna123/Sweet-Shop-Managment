package models

import "gorm.io/gorm"

type Sweet struct {
	gorm.Model
	Name     string  `gorm:"not null" json:"name"`
	Category string  `gorm:"not null" json:"category"`
	Price    float64 `gorm:"not null" json:"price"`
	Quantity int     `gorm:"not null" json:"quantity"`
}
