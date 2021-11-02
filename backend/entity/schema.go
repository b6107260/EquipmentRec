package entity

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Doctor_id          string
	Password           string
	Doctor_name        string
	Pid                string              `gorm:"uniqueIndex"`
	RequisitionRecords []RequisitionRecord `gorm:"foreignKey:DoctorID"`
}
type Equipment struct {
	gorm.Model
	Equipment_id       string
	Equipment_name     string
	Equipment_type     string
	Equipment_cost     float32
	RequisitionRecords []RequisitionRecord `gorm:"foreignKey:EquipmentID"`
}

type Admission struct {
	gorm.Model
	PatientID       string
	Patient_Name    string
	RoomID          string
	Right_Treatment string
	AdmitTime       time.Time

	RequisitionRecords []RequisitionRecord `gorm:"foreignKey:AdmissionID"`
}

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
