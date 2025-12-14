package handlers

import (
	"net/http"

	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/database"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	var transactions []models.Transaction
	// Preload User and Sweet to get details
	if err := database.DB.Preload("User").Preload("Sweet").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
