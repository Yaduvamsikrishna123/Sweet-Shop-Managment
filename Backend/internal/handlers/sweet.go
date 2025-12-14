package handlers

import (
	"net/http"
	"strconv"

	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"github.com/gin-gonic/gin"
)

type SweetInput struct {
	Name     string  `json:"name" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
}

func (h *Handler) AddSweet(c *gin.Context) {
	var input SweetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sweet := models.Sweet{
		Name:     input.Name,
		Category: input.Category,
		Price:    input.Price,
		Quantity: input.Quantity,
	}

	if err := h.SweetRepo.Create(&sweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add sweet"})
		return
	}

	c.JSON(http.StatusCreated, sweet)
}

func (h *Handler) ListSweets(c *gin.Context) {
	sweets, err := h.SweetRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sweets"})
		return
	}

	c.JSON(http.StatusOK, sweets)
}

func (h *Handler) UpdateSweet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input SweetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sweet, err := h.SweetRepo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sweet not found"})
		return
	}

	sweet.Name = input.Name
	sweet.Category = input.Category
	sweet.Price = input.Price
	sweet.Quantity = input.Quantity

	if err := h.SweetRepo.Update(sweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sweet"})
		return
	}

	c.JSON(http.StatusOK, sweet)
}

func (h *Handler) DeleteSweet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.SweetRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sweet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sweet deleted successfully"})
}

func (h *Handler) SearchSweets(c *gin.Context) {
	query := c.Query("q")
	sweets, err := h.SweetRepo.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search sweets"})
		return
	}
	c.JSON(http.StatusOK, sweets)
}
