package entity

import (
	"time"

	"gorm.io/gorm"
)

type RequisitionRecord struct {
	gorm.Model
	RecTime     time.Time
	EquipAmount uint
	// Doctor ทำหน้าที่เป็น FK
	DoctorID *uint
	Doctor   Doctor `gorm:"references:id"`
	// Equipment_ID ทำหน้าที่เป็น FK
	EquipmentID *uint
	EquipCost   float32
	Equipment   Equipment `gorm:"references:id"`

	AdmissionID *uint
	Admission   Admission `gorm:"references:id"`
}
