package controller

import (
	"net/http"

	"github.com/b6107260/sa-64-equipment/entity"
	"github.com/gin-gonic/gin"
)

// POST /admissions
func CreateAdmission(c *gin.Context) {
	var admission entity.Admission
	if err := c.ShouldBindJSON(&admission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&admission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": admission})
}

// GET /admission/:id
func GetAdmission(c *gin.Context) {
	var admission entity.Admission
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM admissions WHERE id = ?", id).Scan(&admission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admission})
}

// GET /admissions
func ListAdmission(c *gin.Context) {
	var admission []entity.Admission
	if err := entity.DB().Raw("SELECT * FROM admissions").Scan(&admission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admission})
}

// DELETE /admissions/:id
func DeleteAdmission(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM admissions WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admission not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /admissions
func UpdateAdmission(c *gin.Context) {
	var admission entity.Admission
	if err := c.ShouldBindJSON(&admission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", admission.ID).First(&admission); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admission not found"})
		return
	}

	if err := entity.DB().Save(&admission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admission})
}