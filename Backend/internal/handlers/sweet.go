package handlers

import (
	"net/http"
	"strconv"

	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/repository"
	"github.com/gin-gonic/gin"
)

func AddSweet(c *gin.Context) {
	var sweet models.Sweet
	if err := c.ShouldBindJSON(&sweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.CreateSweet(&sweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add sweet"})
		return
	}

	c.JSON(http.StatusCreated, sweet)
}

func ListSweets(c *gin.Context) {
	sweets, err := repository.GetAllSweets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sweets"})
		return
	}

	c.JSON(http.StatusOK, sweets)
}

func GetSweet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sweet, err := repository.GetSweetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sweet not found"})
		return
	}

	c.JSON(http.StatusOK, sweet)
}

func UpdateSweet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sweet, err := repository.GetSweetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sweet not found"})
		return
	}

	var input models.Sweet
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sweet.Name = input.Name
	sweet.Category = input.Category
	sweet.Price = input.Price
	sweet.Quantity = input.Quantity

	if err := repository.UpdateSweet(sweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sweet"})
		return
	}

	c.JSON(http.StatusOK, sweet)
}

func DeleteSweet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.DeleteSweet(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sweet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sweet deleted successfully"})
}

func SearchSweets(c *gin.Context) {
	query := c.Query("q")
	sweets, err := repository.SearchSweets(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search sweets"})
		return
	}

	c.JSON(http.StatusOK, sweets)
}
