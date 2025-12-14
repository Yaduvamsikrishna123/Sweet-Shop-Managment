package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	UserID     uint      `gorm:"not null" json:"user_id"`
	User       User      `gorm:"foreignKey:UserID" json:"user"`
	SweetID    uint      `gorm:"not null" json:"sweet_id"`
	Sweet      Sweet     `gorm:"foreignKey:SweetID" json:"sweet"`
	Quantity   int       `gorm:"not null" json:"quantity"`
	TotalPrice float64   `gorm:"not null" json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
