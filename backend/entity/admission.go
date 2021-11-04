package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admission struct {
	gorm.Model
	AdmitTime time.Time

	RightTreatmentID *uint
	RightTreatment   RightTreatment

	RoomID *uint
	Room   Room

	PatientID   *uint
	Patient     Patient
	PatientName string

	TreatmentRecords   []TreatmentRecord   `gorm:"foreignKey:AdmissionID"`
	RequisitionRecords []RequisitionRecord `gorm:"foreignKey:AdmissionID"`
}
