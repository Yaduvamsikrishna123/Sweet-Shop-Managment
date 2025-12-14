package repository

import (
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
