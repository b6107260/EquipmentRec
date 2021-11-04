package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/ProjectG10/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload login body
type LoginPharmacistPayload struct {
	Pid      string `json:"pid"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginPharmacistResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

// POST /login
func LoginPharmacist(c *gin.Context) {
	var payload LoginPharmacistPayload
	var user entity.Pharmacist

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา user ด้วย pid ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM pharmacists WHERE pid = ?", payload.Pid).Scan(&user).Error; err != nil {
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

	tokenResponse := LoginPharmacistResponse{
		Token: signedToken,
		ID:    user.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

// GET /users
// List all users
func ListPharmacists(c *gin.Context) {
	var pharmacists []entity.Pharmacist
	if err := entity.DB().Raw("SELECT * FROM pharmacists").Scan(&pharmacists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pharmacists})
}

// GET /user/:id
// Get user by id
func GetPharmacist(c *gin.Context) {
	var pharmacist entity.Pharmacist
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM pharmacists WHERE id = ?", id).Scan(&pharmacist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pharmacist})
}

// POST /users
func CreatePharmacist(c *gin.Context) {
	var pharmacist entity.Pharmacist
	if err := c.ShouldBindJSON(&pharmacist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(pharmacist.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	pharmacist.Password = string(bytes)

	if err := entity.DB().Create(&pharmacist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pharmacist})
}
