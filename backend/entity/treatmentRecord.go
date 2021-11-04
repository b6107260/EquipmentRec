package entity

import (
	"time"

	"gorm.io/gorm"
)

type TreatmentRecord struct {
	gorm.Model
	RecordDate time.Time
	Treatment  string
	Food_type  string
	Med_amount uint
	Cost       int

	EquipmentID *uint
	Equipment   Equipment `gorm:"references:id"`

	DoctorID *uint
	Doctor   Doctor `gorm:"references:id"`

	MedID *uint
	Med   Medicine

	AdmissionID *uint
	Admission   Admission

	MedicationRecord  []MedicationRecord `gorm:"foreignKey:AdmissionID"`
	MedicationRecords []MedicationRecord `gorm:"foreignKey:TreatmentID"`
	FoodAllocate      []Foodallocate     `gorm:"foreignKey:AdmissionID"`
	Foodallocates     []Foodallocate     `gorm:"foreignKey:TreatmentID"`
	Bills             []Bill             `gorm:"foreignKey:TreatmentRecordID"`
}
