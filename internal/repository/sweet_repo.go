package repository

import (
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/database"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
)

func CreateSweet(sweet *models.Sweet) error {
	return database.DB.Create(sweet).Error
}

func GetAllSweets() ([]models.Sweet, error) {
	var sweets []models.Sweet
	err := database.DB.Find(&sweets).Error
	return sweets, err
}

func GetSweetByID(id uint) (*models.Sweet, error) {
	var sweet models.Sweet
	err := database.DB.First(&sweet, id).Error
	return &sweet, err
}

func UpdateSweet(sweet *models.Sweet) error {
	return database.DB.Save(sweet).Error
}

func DeleteSweet(id uint) error {
	return database.DB.Delete(&models.Sweet{}, id).Error
}

func SearchSweets(query string) ([]models.Sweet, error) {
	var sweets []models.Sweet
	// Simple search by name or category
	err := database.DB.Where("name ILIKE ? OR category ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&sweets).Error
	return sweets, err
}
