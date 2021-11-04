package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pharmacist struct {
	gorm.Model
	Name              string
	Password          string
	Pid               string             `gorm:"uniqueIndex"`
	MedicationRecords []MedicationRecord `gorm:"foreignKey:PharmaID"`
}

type MedicationRecord struct {
	gorm.Model
	Amount     uint
	RecordTime time.Time

	PharmaID *uint
	Pharma   Pharmacist

	MedID *uint
	Med   Medicine

	TreatmentID *uint
	AdmissionID *uint

	Treatment TreatmentRecord
}
