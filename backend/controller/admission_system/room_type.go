package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/gin-gonic/gin"
)

// POST /roomtypes
func CreateRoomtype(c *gin.Context) {
	var roomtype entity.Roomtypes
	if err := c.ShouldBindJSON(&roomtype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}

// GET /roomtype/:id
func GetRoomtype(c *gin.Context) {
	var roomtype entity.Roomtypes
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM roomtypes WHERE id = ?", id).Scan(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}

// GET /roomtypes
func ListRoomtypes(c *gin.Context) {
	var roomtypes []entity.Roomtypes
	if err := entity.DB().Raw("SELECT * FROM roomtypes").Scan(&roomtypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomtypes})
}
