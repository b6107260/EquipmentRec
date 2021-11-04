package entity

import (
	"gorm.io/gorm"
)

// -- นักโภชนาการ --
type Nutritionist struct {
	gorm.Model
	Name          string
	Pid           string `gorm:"uniqueIndex"`
	Password      string
	Foodallocates []Foodallocate `gorm:"foreignKey:NutritionistID"`
}

// -- เซตอาหาร --
type Foodset struct {
	gorm.Model
	Foodmenu      string
	Fooddrink     string
	Setprice      string
	Foodallocates []Foodallocate `gorm:"foreignKey:FoodsetID"`
}

// -- เวลาทานอาหาร --
type Foodtime struct {
	gorm.Model
	Foodtime      string
	Foodallocates []Foodallocate `gorm:"foreignKey:FoodtimeID"`
}

// -- การจัดสรรอาหาร --
type Foodallocate struct {
	gorm.Model

	// NutritionistID ทำหน้าที่เป็น FK
	NutritionistID *uint
	Nutritionist   Nutritionist `gorm:"references:id"`
	// TreatmentRecordID ทำหน้าที่เป็น FK
	TreatmentID *uint
	AdmissionID *uint
	Treatment   TreatmentRecord `gorm:"references:id"`
	// FoodSetID ทำหน้าที่เป็น FK
	FoodsetID *uint
	Foodset   Foodset `gorm:"references:id"`
	// FoodTimeID ทำหน้าที่เป็น FK
	FoodtimeID *uint
	Foodtime   Foodtime `gorm:"references:id"`
}
