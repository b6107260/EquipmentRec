package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/ProjectG10/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload login body
type LoginNursePayload struct {
	Pid      string `json:"pid"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginNurseResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

// POST /login
func LoginNurse(c *gin.Context) {
	var payload LoginNursePayload
	var nurse entity.Nurse

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา nurse ด้วย pid ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM nurses WHERE pid = ?", payload.Pid).Scan(&nurse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(nurse.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid nurse credentials"})
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

	signedToken, err := jwtWrapper.GenerateToken(nurse.Pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginNurseResponse{
		Token: signedToken,
		ID:    nurse.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

// POST /nurse
func CreateNurse(c *gin.Context) {
	var nurse entity.Nurse
	if err := c.ShouldBindJSON(&nurse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(nurse.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	nurse.Password = string(bytes)

	if err := entity.DB().Create(&nurse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": nurse})
}

// GET /nurse/:id
func GetNurse(c *gin.Context) {
	var nurse entity.Nurse
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM nurses WHERE id = ?", id).Scan(&nurse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nurse})
}

// GET /nurses
func ListNurses(c *gin.Context) {
	var nurses []entity.Nurse
	if err := entity.DB().Raw("SELECT * FROM nurses").Scan(&nurses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nurses})
}
