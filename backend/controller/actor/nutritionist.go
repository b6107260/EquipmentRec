package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/ProjectG10/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload login body
type LoginNutritionistPayload struct {
	Pid      string `json:"pid"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginNutritionistResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

// POST /login
func LoginNutritionist(c *gin.Context) {
	var payload LoginNutritionistPayload
	var nutritionist entity.Nutritionist

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา nutritionist ด้วย PersonID ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM nutritionists WHERE pid = ?", payload.Pid).Scan(&nutritionist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(nutritionist.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid nutritionist credentials"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token
	// service.JwtWrapper

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(nutritionist.Pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginNutritionistResponse{
		Token: signedToken,
		ID:    nutritionist.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

// GET /Nutritionist
// List all Nutritionist
func ListNutritionists(c *gin.Context) {
	var nutritionists []entity.Nutritionist
	if err := entity.DB().Raw("SELECT * FROM nutritionists").Scan(&nutritionists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nutritionists})
}

// GET /nutritionists/:id
// Get nutritionist by id
func GetNutritionist(c *gin.Context) {
	var nutritionist entity.Nutritionist
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM nutritionists WHERE id = ?", id).Scan(&nutritionist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nutritionist})
}

// POST /nutritionists
func CreateNutritionist(c *gin.Context) {
	var nutritionist entity.Nutritionist
	if err := c.ShouldBindJSON(&nutritionist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(nutritionist.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	nutritionist.Password = string(bytes)

	if err := entity.DB().Create(&nutritionist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nutritionist})
}

// PATCH /nutritionists
func UpdateNutritionist(c *gin.Context) {
	var nutritionist entity.Nutritionist
	if err := c.ShouldBindJSON(&nutritionist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", nutritionist.ID).First(&nutritionist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nutritionist not found"})
		return
	}

	if err := entity.DB().Save(&nutritionist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nutritionist})
}

// DELETE /nutritionists/:id
func DeleteNutritionist(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM nutritionists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nutritionist not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.Nutritionist{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}
