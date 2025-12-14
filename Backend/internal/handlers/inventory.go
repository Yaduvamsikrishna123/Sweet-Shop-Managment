package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"github.com/gin-gonic/gin"
)

type InventoryInput struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

func (h *Handler) PurchaseSweet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input InventoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sweet, err := h.SweetRepo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sweet not found"})
		return
	}

	if sweet.Quantity < input.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	sweet.Quantity -= input.Quantity
	if err := h.SweetRepo.Update(sweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	// Record Transaction
	var userID uint
	if idVal, exists := c.Get("userId"); exists {
		if idFloat, ok := idVal.(float64); ok {
			userID = uint(idFloat)
		} else {
			log.Printf("userId is of type %T: %v", idVal, idVal)
		}
	}

	transaction := models.Transaction{
		UserID:     userID,
		SweetID:    sweet.ID,
		Quantity:   input.Quantity,
		TotalPrice: float64(input.Quantity) * sweet.Price,
	}

	if err := h.TransactionRepo.Create(&transaction); err != nil {
		log.Printf("Failed to record transaction: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Purchase successful", "remaining_quantity": sweet.Quantity})
}

func (h *Handler) RestockSweet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input InventoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sweet, err := h.SweetRepo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sweet not found"})
		return
	}

	sweet.Quantity += input.Quantity
	if err := h.SweetRepo.Update(sweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Restock successful", "new_quantity": sweet.Quantity})
}
