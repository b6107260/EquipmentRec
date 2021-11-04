package controller

import (
	"net/http"

	"github.com/ProjectG10/entity"
	"github.com/gin-gonic/gin"
)

// POST /food_allocate
func CreateFoodallocate(c *gin.Context) {

	var foodallocate entity.Foodallocate
	var treatmentrecord entity.TreatmentRecord
	var foodset entity.Foodset
	var foodtime entity.Foodtime
	var nutritionist entity.Nutritionist

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร foodallocate
	if err := c.ShouldBindJSON(&foodallocate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา treatmentrecord ด้วย id
	if tx := entity.DB().Where("id = ?", foodallocate.TreatmentID).First(&treatmentrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	// 11: ค้นหา foodset ด้วย id
	if tx := entity.DB().Where("id = ?", foodallocate.FoodsetID).First(&foodset); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	// 12: ค้นหา foodtime ด้วย id
	if tx := entity.DB().Where("id = ?", foodallocate.FoodtimeID).First(&foodtime); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}
	// 13: ค้นหา nutritionist ด้วย id
	if tx := entity.DB().Where("id = ?", foodallocate.NutritionistID).First(&nutritionist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}

	// 12: สร้าง foodallocate
	fa := entity.Foodallocate{
		Treatment:    treatmentrecord, // โยงความสัมพันธ์กับ Entity TreatmentRecord
		Foodset:      foodset,         // โยงความสัมพันธ์กับ Entity FoodSet
		Foodtime:     foodtime,        // โยงความสัมพันธ์กับ Entity FoodTime
		Nutritionist: nutritionist,    // โยงความสัมพันธ์กับ Entity nutritionist
		AdmissionID:  treatmentrecord.AdmissionID,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&fa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": fa})
}

// GET /foodallocate/:id
func GetFoodallocate(c *gin.Context) {
	var foodallocate entity.Foodallocate
	id := c.Param("id")
	if err := entity.DB().Preload("Treatmentrecord").Preload("Foodset").Preload("Foodtime").Preload("Nutritionist").Raw("SELECT * FROM foodallocates WHERE id = ?", id).Find(&foodallocate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foodallocate})
}

// GET /food_allocate
func ListFoodallocates(c *gin.Context) {
	var foodallocates []entity.Foodallocate
	if err := entity.DB().Preload("Treatmentrecord").Preload("Foodset").Preload("Foodtime").Preload("Nutritionist").Raw("SELECT * FROM foodallocates").Find(&foodallocates).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodallocates})
}

// DELETE /watch_videos/:id
func DeleteFoodallocate(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM foodallocates WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "foodallocate not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_videos
func UpdateFoodallocate(c *gin.Context) {
	var foodallocate entity.Foodallocate
	if err := c.ShouldBindJSON(&foodallocate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", foodallocate.ID).First(&foodallocate); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "foodallocate not found"})
		return
	}

	if err := entity.DB().Save(&foodallocate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodallocate})
}
