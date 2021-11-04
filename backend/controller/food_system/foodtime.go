package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/gin-gonic/gin"
)

// POST /foodtime
func CreateFoodTime(c *gin.Context) {
	var foodtime entity.Foodtime
	if err := c.ShouldBindJSON(&foodtime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&foodtime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foodtime})
}

// GET /foodtime/:id
func GetFoodTime(c *gin.Context) {
	var foodtime entity.Foodtime

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM foodtimes WHERE id = ?", id).Find(&foodtime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodtime})
}

// GET /foodtimes
func ListFoodTimes(c *gin.Context) {
	var foodtimes []entity.Foodtime
	if err := entity.DB().Raw("SELECT * FROM foodtimes").Find(&foodtimes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodtimes})
}

// DELETE /foodtimes/:id
func DeleteFoodTime(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM foodtimes WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "foodtime not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /foodtimes
func UpdateFoodTime(c *gin.Context) {
	var foodtime entity.Foodtime
	if err := c.ShouldBindJSON(&foodtime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", foodtime.ID).First(&foodtime); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "foodtime not found"})
		return
	}

	if err := entity.DB().Save(&foodtime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodtime})
}
