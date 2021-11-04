package entity

import (
	"gorm.io/gorm"
)

type RightTreatment struct {
	gorm.Model

	RightTreatmentName   string `gorm:"uniqueIndex"`
	RightTreatmentDetail string
	Price                int

	Admissions []Admission `gorm:"foreignKey:RightTreatmentID"`
	Bills      []Bill      `gorm:"foreignKey:RightTreatmentID"`
}
