package repository

import "github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"

type SweetRepository interface {
	Create(sweet *models.Sweet) error
	GetAll() ([]models.Sweet, error)
	GetByID(id uint) (*models.Sweet, error)
	Update(sweet *models.Sweet) error
	Delete(id uint) error
	Search(query string) ([]models.Sweet, error)
}

type UserRepository interface {
	Create(user *models.User) error
	GetByUsername(username string) (*models.User, error)
}

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	GetAll() ([]models.Transaction, error)
}
