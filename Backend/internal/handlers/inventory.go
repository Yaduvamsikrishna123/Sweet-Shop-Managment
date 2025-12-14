package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/database"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/repository"
	"github.com/gin-gonic/gin"
)

type InventoryInput struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

func PurchaseSweet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input InventoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sweet, err := repository.GetSweetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sweet not found"})
		return
	}

	if sweet.Quantity < input.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient quantity"})
		return
	}

	sweet.Quantity -= input.Quantity
	if err := repository.UpdateSweet(sweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	// Record Transaction
	userID := c.GetUint("userID") // Assuming AuthMiddleware sets this
	transaction := models.Transaction{
		UserID:     userID,
		SweetID:    sweet.ID,
		Quantity:   input.Quantity,
		TotalPrice: float64(input.Quantity) * sweet.Price,
	}
	if err := database.DB.Create(&transaction).Error; err != nil {
		// Log error but don't fail the request as purchase was successful
		// In a real app, we might want to use a transaction block
		log.Printf("Failed to record transaction: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Purchase successful", "remaining_quantity": sweet.Quantity})
}

func RestockSweet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input InventoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sweet, err := repository.GetSweetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sweet not found"})
		return
	}

	sweet.Quantity += input.Quantity
	if err := repository.UpdateSweet(sweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Restock successful", "new_quantity": sweet.Quantity})
}
