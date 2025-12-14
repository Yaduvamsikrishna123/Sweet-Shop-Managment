package repository

import (
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"gorm.io/gorm"
)

type GormSweetRepository struct {
	DB *gorm.DB
}

func NewSweetRepository(db *gorm.DB) SweetRepository {
	return &GormSweetRepository{DB: db}
}

func (r *GormSweetRepository) Create(sweet *models.Sweet) error {
	return r.DB.Create(sweet).Error
}

func (r *GormSweetRepository) GetAll() ([]models.Sweet, error) {
	var sweets []models.Sweet
	err := r.DB.Find(&sweets).Error
	return sweets, err
}

func (r *GormSweetRepository) GetByID(id uint) (*models.Sweet, error) {
	var sweet models.Sweet
	err := r.DB.First(&sweet, id).Error
	return &sweet, err
}

func (r *GormSweetRepository) Update(sweet *models.Sweet) error {
	return r.DB.Save(sweet).Error
}

func (r *GormSweetRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Sweet{}, id).Error
}

func (r *GormSweetRepository) Search(query string) ([]models.Sweet, error) {
	var sweets []models.Sweet
	err := r.DB.Where("name ILIKE ? OR category ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&sweets).Error
	return sweets, err
}
