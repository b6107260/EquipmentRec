package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/gin-gonic/gin"
)

// POST /medication_records
func CreateMedicationRecord(c *gin.Context) {

	var medicationrecord entity.MedicationRecord
	var pharmacist entity.Pharmacist
	var treatmentrecord entity.TreatmentRecord
	var medicine entity.Medicine

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร medicationrecord
	if err := c.ShouldBindJSON(&medicationrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา pharmacist ด้วย id
	if tx := entity.DB().Where("id = ?", medicationrecord.PharmaID).First(&pharmacist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pharmacist not found"})
		return
	}

	// 9: ค้นหา treatmentrecord ด้วย id
	if tx := entity.DB().Where("id = ?", medicationrecord.TreatmentID).First(&treatmentrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "treatment record not found"})
		return
	}

	// 10: ค้นหา medicine ด้วย id
	if tx := entity.DB().Where("id = ?", medicationrecord.MedID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}

	// 11: สร้าง MedicationRecord
	mr := entity.MedicationRecord{
		Pharma:      pharmacist,                  // โยงความสัมพันธ์กับ Entity Pharmacist
		Med:         medicine,                    // โยงความสัมพันธ์กับ Entity Medicine
		Treatment:   treatmentrecord,             // โยงความสัมพันธ์กับ Entity TreatmentRecord
		Amount:      medicationrecord.Amount,     // ตั้งค่า Amount
		RecordTime:  medicationrecord.RecordTime, // ตั้งค่า RecordTime
		AdmissionID: treatmentrecord.AdmissionID,
	}

	// 12: บันทึก
	if err := entity.DB().Create(&mr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mr})
}

// GET /medication_record/:id
func GetMedicationRacord(c *gin.Context) {
	var medicationrecord entity.MedicationRecord
	id := c.Param("id")
	if err := entity.DB().Preload("Treatment").Preload("Med").Preload("Pharma").Raw("SELECT * FROM medication_records WHERE id = ?", id).Find(&medicationrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicationrecord})
}

// GET /medication_records
func ListMedicationRacord(c *gin.Context) {
	var medicationrecord []entity.MedicationRecord
	if err := entity.DB().Preload("Treatment").Preload("Med").Preload("Pharma").Preload("Treatment.Admission").Raw("SELECT * FROM medication_records").Find(&medicationrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicationrecord})
}
