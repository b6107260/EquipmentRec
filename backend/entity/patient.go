package entity

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	Date         time.Time
	Patient_name string

	Identification string `gorm:"unqueIndex"`

	Record_byID *uint
	Record_by   Nurse

	DiseaseID *uint
	Disease   Disease

	MedID *uint
	Med   Medicine

	Admissions []Admission `gorm:"foreignKey:PatientID"`
}
