package entity

import (
	"gorm.io/gorm"
)

type Roomtypes struct {
	gorm.Model
	Name  string
	Price int

	Rooms []Room `gorm:"foreignKey:RoomtypesID"`
}

type Room struct {
	gorm.Model
	Number string

	RoomtypesID *uint
	Roomtypes   Roomtypes

	Admissions []Admission `gorm:"foreignKey:RoomID"`
}
