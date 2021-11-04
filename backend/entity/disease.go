package entity

import (
	"gorm.io/gorm"
)

type Disease struct {
	gorm.Model
	Disease_name string    `gorm:"uniqueIndex"`
	Patients     []Patient `gorm:"foreignKey:DiseaseID"`
}
