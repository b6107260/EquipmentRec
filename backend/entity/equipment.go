package entity

import (
	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	Equipment_id       string
	Equipment_name     string
	Equipment_type     string
	Equipment_cost     float32
	RequisitionRecords []RequisitionRecord `gorm:"foreignKey:EquipmentID"`
	TreatmentRecords   []TreatmentRecord   `gorm:"foreignKey:EquipmentID"`
}
