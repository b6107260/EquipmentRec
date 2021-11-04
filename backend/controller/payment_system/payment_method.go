package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/gin-gonic/gin"
)

// GET /PaymentMethod
// List all PaymentMethod
func ListPaymentMethod(c *gin.Context) {
	var paymentmethod []entity.PaymentMethod
	if err := entity.DB().Raw("SELECT * FROM payment_methods").Scan(&paymentmethod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentmethod})
}

// GET /PaymentMethod/:id
// Get PaymentMethod by id
func GetPaymentMethod(c *gin.Context) {
	var paymentmethod entity.PaymentMethod
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM payment_methods WHERE id = ?", id).Scan(&paymentmethod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentmethod})
}

// POST /PaymentMethod
func CreatePaymentMethod(c *gin.Context) {
	var paymentmethod entity.PaymentMethod
	if err := c.ShouldBindJSON(&paymentmethod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&paymentmethod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentmethod})
}
