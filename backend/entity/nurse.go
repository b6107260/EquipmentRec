package entity

import (
	"gorm.io/gorm"
)

type Nurse struct {
	gorm.Model
	Pid        string `gorm:"uniqueIndex"`
	Nurse_name string
	Password   string
	Tel        string `gorm:"uniqueIndex"`

	Patients []Patient `gorm:"foreignKey:Record_byID"`
}
