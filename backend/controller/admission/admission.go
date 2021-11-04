package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/gin-gonic/gin"
)

func CreateAdmission(c *gin.Context) {

	var admission entity.Admission
	var patient entity.Patient
	var room entity.Room
	//var roomtypes entity.Roomtypes
	var righttreatment entity.RightTreatment

	// ผลลัพธ์ที่ได้จะถูก bind เข้าตัวแปร Admission
	if err := c.ShouldBindJSON(&admission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา patient ด้วย id
	if tx := entity.DB().Where("id = ?", admission.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	// ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", admission.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	// ค้นหา right_treatment ด้วย id
	if tx := entity.DB().Where("id = ?", admission.RightTreatmentID).First(&righttreatment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "right_treatment not found"})
		return
	}

	// สร้าง admission
	ad := entity.Admission{
		RightTreatment: righttreatment,      // โยงความสัมพันธ์กับ Entity Right_Treatment
		Room:           room,                // โยงความสัมพันธ์กับ Entity Room
		Patient:        patient,             // โยงความสัมพันธ์กับ Entity Patient
		AdmitTime:      admission.AdmitTime, // ตั้งค่าฟิลด์ AdmitTime
	}

	// บันทึก
	if err := entity.DB().Create(&ad).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ad})
}

// GET /Admission
func ListAdmission(c *gin.Context) {
	var admission []entity.Admission
	if err := entity.DB().Preload("Room").Preload("Patient").Preload("RightTreatment").Raw("SELECT * FROM admissions").Find(&admission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admission})
}

// GET /admission/:id
func GetAdmission(c *gin.Context) {
	var admission entity.Admission
	id := c.Param("id")
	if err := entity.DB().Preload("Room").Preload("Patient").Preload("RightTreatment").Raw("SELECT * FROM admissions WHERE id = ?", id).Find(&admission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": admission})
}
