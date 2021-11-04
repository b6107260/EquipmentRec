package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/gin-gonic/gin"
)

// POST /requisition_records
func CreateRequisitionRecord(c *gin.Context) {

	var requisitionrecord entity.RequisitionRecord
	var doctor entity.Doctor
	var admission entity.Admission
	var equipment entity.Equipment

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร requisitionrecord
	if err := c.ShouldBindJSON(&requisitionrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา doctor ด้วย id
	if tx := entity.DB().Where("id = ?", requisitionrecord.DoctorID).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	// 9: ค้นหา admission ด้วย id
	if tx := entity.DB().Where("id = ?", requisitionrecord.AdmissionID).First(&admission); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admission record not found"})
		return
	}

	// 10: ค้นหา Equipment ด้วย id
	if tx := entity.DB().Where("id = ?", requisitionrecord.EquipmentID).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}

	// 11: สร้าง RequisitionRecord
	rc := entity.RequisitionRecord{
		Doctor:      doctor,
		Equipment:   equipment,
		Admission:   admission,
		EquipAmount: requisitionrecord.EquipAmount,
		RecTime:     requisitionrecord.RecTime,
		EquipCost:   equipment.Equipment_cost,
	}
	// 12: บันทึก
	if err := entity.DB().Create(&rc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rc})
}

// GET /requisition_record/:id
func GetRequisitionRecord(c *gin.Context) {
	var requisitionrecord entity.RequisitionRecord
	id := c.Param("id")
	if err := entity.DB().Preload("Admission").Preload("Equipment").Preload("Doctor").Raw("SELECT * FROM requisition_records WHERE id = ?", id).Find(&requisitionrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": requisitionrecord})
}

// GET /requisition_records
func ListRequisitionRecord(c *gin.Context) {
	var requisitionrecord []entity.RequisitionRecord
	if err := entity.DB().Preload("Admission").Preload("Equipment").Preload("Doctor").Raw("SELECT * FROM requisition_records").Find(&requisitionrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requisitionrecord})
}

// DELETE /requisition_records/:id
func DeleteRequisitionRecord(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM requisition_record WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "requisition record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /requisition_records
func UpdateRequisitionRecord(c *gin.Context) {
	var requisitionrecord entity.RequisitionRecord
	if err := c.ShouldBindJSON(&requisitionrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", requisitionrecord.ID).First(&requisitionrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "requisition record not found"})
		return
	}

	if err := entity.DB().Save(&requisitionrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requisitionrecord})
}
