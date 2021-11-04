package entity

import (
	"time"

	"gorm.io/gorm"
)

type FinancialOfficer struct {
	gorm.Model
	FinancialName string
	Pid           string `gorm:"uniqueIndex"`
	Password      string
}

type PaymentMethod struct {
	gorm.Model
	PaymentMethodName string `gorm:"uniqueIndex"`

	Bills []Bill `gorm:"foreignKey:PaymentMethodID"`
}

type Bill struct {
	gorm.Model
	BillDateTime      time.Time
	TreatmentRecordID *uint
	TreatmentRecord   TreatmentRecord
	RightTreatmentID  *uint
	RightTreatment    RightTreatment
	PaymentMethodID   *uint
	PaymentMethod     PaymentMethod
	AmountPaid        int
}
