package repository

import (
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"gorm.io/gorm"
)

type GormTransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &GormTransactionRepository{DB: db}
}

func (r *GormTransactionRepository) Create(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}

func (r *GormTransactionRepository) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Preload("User").Preload("Sweet").Find(&transactions).Error
	return transactions, err
}
