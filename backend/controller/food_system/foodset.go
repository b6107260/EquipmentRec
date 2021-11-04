package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/gin-gonic/gin"
)

// POST /foodset
func CreateFoodSet(c *gin.Context) {
	var foodset entity.Foodset
	if err := c.ShouldBindJSON(&foodset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&foodset).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foodset})
}

// GET /foodset/:id
func GetFoodSet(c *gin.Context) {
	var foodset entity.Foodset

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM foodsets WHERE id = ?", id).Find(&foodset).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodset})
}

// GET /foodsets
func ListFoodSets(c *gin.Context) {
	var foodsets []entity.Foodset
	if err := entity.DB().Raw("SELECT * FROM foodsets").Find(&foodsets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodsets})
}

// DELETE /tfoodsets/:id
func DeleteFoodSet(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM foodsets WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "foodset not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /FoodSets
func UpdateFoodSet(c *gin.Context) {
	var foodset entity.Foodset
	if err := c.ShouldBindJSON(&foodset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", foodset.ID).First(&foodset); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "foodset not found"})
		return
	}

	if err := entity.DB().Save(&foodset).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodset})
}
