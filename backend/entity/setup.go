package entity

import (
	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	//Migrate the schema

	database.AutoMigrate(
		&Doctor{},
		&Equipment{},
		&Admission{},
		&RequisitionRecord{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&Doctor{}).Create(&Doctor{
		Doctor_id:   "5001",
		Doctor_name: "Dr.Sonsak",
		Pid:  "123456789994",
		Password:    string(password),
	})

	db.Model(&Doctor{}).Create(&Doctor{
		Doctor_id:   "5002",
		Doctor_name: "Dr.Narumon",
		Pid:  "2222245665666",
		Password:    string(password),
	})

	var songsak Doctor
	var narumon Doctor

	db.Raw("SELECT * FROM doctors WHERE doctor_id = ?", "5001").Scan(&songsak)
	db.Raw("SELECT * FROM doctors WHERE doctor_id = ?", "5002").Scan(&narumon)

	db.Model(&Equipment{}).Create(&Equipment{
		Equipment_id:   "001",
		Equipment_name: "Syringe",
		Equipment_type: "surgical equipment",
		Equipment_cost: 1452.50,
	})

	db.Model(&Equipment{}).Create(&Equipment{
		Equipment_id:   "002",
		Equipment_name: "Surgical Blade",
		Equipment_type: "surgical equipment",
		Equipment_cost: 150.00,
	})

	var syringe Equipment
	var blade Equipment

	db.Raw("SELECT * FROM equipment WHERE equipment_id = ?", "001").Scan(&syringe)
	db.Raw("SELECT * FROM equipment WHERE equipment_id = ?", "002").Scan(&blade)

	db.Model(&Admission{}).Create(&Admission{
		PatientID:       "A001",
		Patient_Name:    "Janpen Thongkham",
		RoomID:          "RM2005",
		Right_Treatment: "IV0003",
	})

	db.Model(&Admission{}).Create(&Admission{
		PatientID:       "A002",
		Patient_Name:    "Nattaporn Bud",
		RoomID:          "RM2004",
		Right_Treatment: "IV0002",
	})
	var janpen Admission
	var nattaporn Admission

	db.Raw("SELECT * FROM admissions WHERE patient_id = ?", "A001").Scan(&janpen)
	db.Raw("SELECT * FROM admissions WHERE patient_id = ?", "A002").Scan(&nattaporn)

}
