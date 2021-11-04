package entity

import (
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Password           string
	Doctor_name        string
	Pid                string              `gorm:"uniqueIndex"`
	RequisitionRecords []RequisitionRecord `gorm:"foreignKey:DoctorID"`
}
