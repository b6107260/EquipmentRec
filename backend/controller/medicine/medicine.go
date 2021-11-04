package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/gin-gonic/gin"
)

// POST /medicine
func CreateMedicine(c *gin.Context) {
	var medicine entity.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&medicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicine})
}

// GET /medicine/:id
func GetMedicine(c *gin.Context) {
	var medicine entity.Medicine
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medicines WHERE id = ?", id).Scan(&medicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicine})
}

// GET /medicines
func ListMedicine(c *gin.Context) {
	var medicine []entity.Medicine
	if err := entity.DB().Raw("SELECT * FROM medicines").Scan(&medicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicine})
}
