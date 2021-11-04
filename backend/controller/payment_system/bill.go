package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/gin-gonic/gin"
)

func CreateBill(c *gin.Context) {

	var bill entity.Bill
	var TreatmentRecord entity.TreatmentRecord
	//var patient entity.Patient
	var RightTreatment entity.RightTreatment
	var PaymentMethod entity.PaymentMethod

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร medicationrecord
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 8: ค้นหา TreatmentRecord ด้วย id
	if tx := entity.DB().Where("id = ?", bill.TreatmentRecordID).First(&TreatmentRecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TreatmentRecord not found"})
		return
	}

	// 10: ค้นหา patient ด้วย id
	/*
		if tx := entity.DB().Where("id = ?", bill.PatientID).First(&patient); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "patient record not found"})
			return
		}*/

	// 11: ค้นหา RightTreatment ด้วย id
	if tx := entity.DB().Where("id = ?", bill.RightTreatmentID).First(&RightTreatment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RightTreatment not found"})
		return
	}
	// 12: ค้นหา PaymentMethod ด้วย id
	if tx := entity.DB().Where("id = ?", bill.PaymentMethodID).First(&PaymentMethod); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentMethod not found"})
		return
	}

	a := 0
	if RightTreatment.RightTreatmentName == "บัตร30" {
		a = 30
	} else if TreatmentRecord.Cost > (-RightTreatment.Price) {
		a = TreatmentRecord.Cost + RightTreatment.Price
	} else {
		a = 0
	}

	//a = (TreatmentRecord.Cost > RightTreatment.Price) ? (TreatmentRecord.Cost - RightTreatment.Price) : ""

	// 14: สร้าง bill
	b := entity.Bill{
		BillDateTime:    bill.BillDateTime,
		TreatmentRecord: TreatmentRecord,
		//Patient:          patient,
		//PatientPID:     TreatmentRecord.Admission.Patient.IdentificationID,
		RightTreatment: RightTreatment,
		PaymentMethod:  PaymentMethod,
		AmountPaid:     a,
	}

	// 15: บันทึก
	if err := entity.DB().Create(&b).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": b})
}

// GET /Bill
// List all Bill
func ListBill(c *gin.Context) {
	var bill []entity.Bill
	if err := entity.DB().Preload("PaymentMethod").Preload("RightTreatment").Preload("TreatmentRecord").Preload("TreatmentRecord.Admission.Patient").Raw("SELECT * FROM bills").Find(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bill})
}

// GET /Bill/:id
// Get Bill by id
func GetBill(c *gin.Context) {
	var bill entity.Bill
	id := c.Param("id")
	if err := entity.DB().Preload("PaymentMethod").Preload("RightTreatment").Preload("TreatmentRecord").Raw("SELECT * FROM bills WHERE id = ?", id).Find(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bill})
}
