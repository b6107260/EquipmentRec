package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/ProjectG10/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload login body
type LoginFinancialOfficerPayload struct {
	Pid      string `json:"pid"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginFinancialOfficerResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

// POST /login
func LoginFinancialOfficer(c *gin.Context) {
	var payload LoginFinancialOfficerPayload
	var user entity.FinancialOfficer

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา user ด้วย pid ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM financial_officers WHERE pid = ?", payload.Pid).Scan(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user credentials"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginFinancialOfficerResponse{
		Token: signedToken,
		ID:    user.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

// GET /FinancialOfficer
// List all FinancialOfficer
func ListFinancialOfficer(c *gin.Context) {
	var financialofficer []entity.FinancialOfficer
	if err := entity.DB().Raw("SELECT * FROM financial_officers").Scan(&financialofficer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": financialofficer})
}

// GET /FinancialOfficer/:id
// Get FinancialOfficer by id
func GetFinancialOfficer(c *gin.Context) {
	var financialofficer entity.FinancialOfficer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM financial_officers WHERE id = ?", id).Scan(&financialofficer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": financialofficer})
}

// POST /FinancialOfficer
func CreateFinancialOfficer(c *gin.Context) {
	var financialofficer entity.FinancialOfficer
	if err := c.ShouldBindJSON(&financialofficer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&financialofficer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": financialofficer})
}
