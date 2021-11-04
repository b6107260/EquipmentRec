package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("g10-db.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	database.AutoMigrate(

		&Admission{}, &Disease{}, &Doctor{}, &Equipment{}, &Foodallocate{}, &Pharmacist{}, &MedicationRecord{},
		&Medicine{}, &Nurse{}, &Patient{}, &FinancialOfficer{}, &PaymentMethod{}, &Bill{},
		&RequisitionRecord{}, &RightTreatment{}, &Roomtypes{}, &Room{}, &TreatmentRecord{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	//-------------Nurse----------------------
	nurse1 := Nurse{
		Nurse_name: "Sumaree Namsommut",
		Tel:        "0988888888",
		Pid:        "1155523456789",
		Password:   string(password),
	}
	db.Model(&Nurse{}).Create(&nurse1)
	nurse2 := Nurse{
		Nurse_name: "Nattawee Namsommut",
		Tel:        "0866663333",
		Pid:        "1166678901234",
		Password:   string(password),
	}
	db.Model(&Nurse{}).Create(&nurse2)

	//----------------FinancialOfficer----------------
	fn1 := FinancialOfficer{
		FinancialName: "Somchai Namsommut",
		Pid:           "1234567890123",
		Password:      string(password),
	}
	db.Model(&FinancialOfficer{}).Create(&fn1)
	fn2 := FinancialOfficer{
		FinancialName: "Samhon Namsommut",
		Pid:           "3210987654321",
		Password:      string(password),
	}
	db.Model(&FinancialOfficer{}).Create(&fn2)

	//-----------------Doctor-----------------
	doctor1 := Doctor{
		Doctor_name: "Dr.Sonsak Namsommut",
		Pid:         "123456789994",
		Password:    string(password),
	}
	db.Model(&Doctor{}).Create(&doctor1)
	doctor2 := Doctor{
		Doctor_name: "Dr.Narumon Namsommut",
		Pid:         "2222245665666",
		Password:    string(password),
	}
	db.Model(&Doctor{}).Create(&doctor2)

	//---------------Nutritionist--------------
	n1 := Nutritionist{
		Name:     "Kita Namsommut",
		Pid:      "0000000000001",
		Password: string(password),
	}
	db.Model(&Nutritionist{}).Create(&n1)

	n2 := Nutritionist{
		Name:     "Run Namsommut",
		Pid:      "0000000000002",
		Password: string(password),
	}
	db.Model(&Nutritionist{}).Create(&n2)

	//--------------Pharmacist-------------
	pharma1 := Pharmacist{
		Name:     "Chawarat Narit",
		Pid:      "1400011111111",
		Password: string(password),
	}
	db.Model(&Pharmacist{}).Create(&pharma1)
	pharma2 := Pharmacist{
		Name:     "Pichanon  Srisongmuang",
		Pid:      "1400011111112",
		Password: string(password),
	}
	db.Model(&Pharmacist{}).Create(&pharma2)
	pharma3 := Pharmacist{
		Name:     "Chin Love",
		Pid:      "1400011111113",
		Password: string(password),
	}
	db.Model(&Pharmacist{}).Create(&pharma3)

	//-------------Disease--------------------
	disease0 := Disease{
		Disease_name: "-",
	}
	db.Model(&Disease{}).Create(&disease0)
	disease1 := Disease{
		Disease_name: "Diabetes",
	}
	db.Model(&Disease{}).Create(&disease1)
	disease2 := Disease{
		Disease_name: "Cancer",
	}
	db.Model(&Disease{}).Create(&disease2)
	disease3 := Disease{
		Disease_name: "Heart Disease",
	}
	db.Model(&Disease{}).Create(&disease3)
	disease4 := Disease{
		Disease_name: "Stroke",
	}
	db.Model(&Disease{}).Create(&disease4)

	//-----------------Roomtypes-------------------
	all := Roomtypes{
		Name:  "ห้องรวม",
		Price: 500,
	}
	db.Model(&Roomtypes{}).Create(&all)

	std := Roomtypes{
		Name:  "Standard Room",
		Price: 1500,
	}
	db.Model(&Roomtypes{}).Create(&std)

	pm := Roomtypes{
		Name:  "Premier Room",
		Price: 2500,
	}
	db.Model(&Roomtypes{}).Create(&pm)

	np := Roomtypes{
		Name:  "Negative pressure",
		Price: 4000,
	}
	db.Model(&Roomtypes{}).Create(&np)

	//---------------------Room------------------
	A10001 := Room{
		Number:    "10001",
		Roomtypes: all,
	}
	db.Model(&Room{}).Create(&A10001)

	std20001 := Room{
		Number:    "STD-20001",
		Roomtypes: std,
	}
	db.Model(&Room{}).Create(&std20001)

	std20002 := Room{
		Number:    "STD-20002",
		Roomtypes: std,
	}
	db.Model(&Room{}).Create(&std20002)

	pm30001 := Room{
		Number:    "PM-30001",
		Roomtypes: pm,
	}
	db.Model(&Room{}).Create(&pm30001)

	np40001 := Room{
		Number:    "NP-40001",
		Roomtypes: np,
	}
	db.Model(&Room{}).Create(&np40001)

	//--------------RightTreatment---------------------------
	NO0001 := RightTreatment{
		RightTreatmentName:   "ไม่ใช้สิทธิการรักษา",
		RightTreatmentDetail: "ไม่มีส่วนลด",
		Price:                0,
	}
	db.Model(&RightTreatment{}).Create(&NO0001)

	GM0001 := RightTreatment{
		RightTreatmentName:   "บัตร30",
		RightTreatmentDetail: "ลดเหลือ 30 บาท",
		Price:                30,
	}
	db.Model(&RightTreatment{}).Create(&GM0001)

	GM0002 := RightTreatment{
		RightTreatmentName:   "ข้าราชการ",
		RightTreatmentDetail: "ลดสูงสุด 20000 บาท",
		Price:                -20000,
	}
	db.Model(&RightTreatment{}).Create(&GM0002)

	IV0001 := RightTreatment{
		RightTreatmentName:   "ประกันชั้น1",
		RightTreatmentDetail: "ลดสูงสุด 30000 บาท",
		Price:                -30000,
	}
	db.Model(&RightTreatment{}).Create(&IV0001)

	IV0002 := RightTreatment{
		RightTreatmentName:   "ประกันชั้น2",
		RightTreatmentDetail: "ลดสูงสุด 15000 บาท",
		Price:                -15000,
	}
	db.Model(&RightTreatment{}).Create(&IV0002)

	IV0003 := RightTreatment{
		RightTreatmentName:   "ประกันชั้น3",
		RightTreatmentDetail: "ลดสูงสุด 7000 บาท",
		Price:                -7000,
	}
	db.Model(&RightTreatment{}).Create(&IV0003)

	//-----------PaymentMethod---------------
	pm1 := PaymentMethod{
		PaymentMethodName: "เงินสด",
	}
	db.Model(&PaymentMethod{}).Create(&pm1)
	pm2 := PaymentMethod{
		PaymentMethodName: "บัตรเครดิต",
	}
	db.Model(&PaymentMethod{}).Create(&pm2)
	pm3 := PaymentMethod{
		PaymentMethodName: "ออนไลน์",
	}
	db.Model(&PaymentMethod{}).Create(&pm3)

	//-------------Equipment----------------
	equip1 := Equipment{
		Equipment_id:   "001",
		Equipment_name: "Syringe",
		Equipment_type: "surgical equipment",
		Equipment_cost: 1452.50,
	}
	db.Model(&Equipment{}).Create(&equip1)
	equip2 := Equipment{
		Equipment_id:   "002",
		Equipment_name: "Surgical Blade",
		Equipment_type: "surgical equipment",
		Equipment_cost: 150.00,
	}
	db.Model(&Equipment{}).Create(&equip2)

	//--------------FoodSet------------------
	set1 := Foodset{
		Foodmenu:  "ผัดกะเพรา",
		Fooddrink: "น้ำผลไม้",
		Setprice:  "50",
	}
	db.Model(&Foodset{}).Create(&set1)
	set2 := Foodset{
		Foodmenu:  "ข้าวต้ม",
		Fooddrink: "น้ำผลไม้",
		Setprice:  "45",
	}
	db.Model(&Foodset{}).Create(&set2)
	set3 := Foodset{
		Foodmenu:  "งดอหาร",
		Fooddrink: "นม",
		Setprice:  "10",
	}
	db.Model(&Foodset{}).Create(&set3)

	//----------------FoodTime----------------
	time1 := Foodtime{
		Foodtime: "เช้า เที่ยง เย็น",
	}
	db.Model(&Foodtime{}).Create(&time1)
	time2 := Foodtime{
		Foodtime: "เช้า เที่ยง",
	}
	db.Model(&Foodtime{}).Create(&time2)
	time3 := Foodtime{
		Foodtime: "เช้า",
	}
	db.Model(&Foodtime{}).Create(&time3)
	time4 := Foodtime{
		Foodtime: "งดอาหาร",
	}
	db.Model(&Foodtime{}).Create(&time4)

	//-------------Medicine------------------
	med1 := Medicine{
		Med_name:  "PARACETAMOL 500 MG",
		Med_type:  "TAB",
		Med_price: 1,
	}
	db.Model(&Medicine{}).Create(&med1)
	med2 := Medicine{
		Med_name:  "CEFCINIR 100 MG",
		Med_type:  "CAP",
		Med_price: 14,
	}
	db.Model(&Medicine{}).Create(&med2)
	med3 := Medicine{
		Med_name:  "PREDNISLONE 5 MG",
		Med_type:  "TAB",
		Med_price: 1,
	}
	db.Model(&Medicine{}).Create(&med3)

	//-----------------------Every Table----------------------------
	//1.-----------------Patient---------------------
	pt1 := Patient{
		Patient_name:   "Som-A Tester",
		Identification: "12xxxxxxxxxxx",
		Disease:        disease0,
		Date:           time.Now(),
		Record_by:      nurse1,
		Med:            med1,
	}
	db.Model(&Patient{}).Create(&pt1)
	pt2 := Patient{
		Patient_name:   "Som-B Tester",
		Identification: "56xxxxxxxxxxx",
		Disease:        disease1,
		Date:           time.Now(),
		Record_by:      nurse2,
		Med:            med2,
	}
	db.Model(&Patient{}).Create(&pt2)
	pt3 := Patient{
		Patient_name:   "Som-C Tester",
		Identification: "79xxxxxxxxxxx",
		Date:           time.Now(),
		Disease:        disease3,
		Record_by:      nurse2,
		Med:            med3,
	}
	db.Model(&Patient{}).Create(&pt3)

	//2.-----------------Admission---------------------
	//Admission
	ad1 := Admission{
		AdmitTime:      time.Now(),
		RightTreatment: GM0001,
		Room:           std20002,
		Patient:        pt1,
		PatientName:    pt1.Patient_name,
	}
	db.Model(&Admission{}).Create(&ad1)

	ad2 := Admission{
		AdmitTime:      time.Now(),
		RightTreatment: IV0003,
		Room:           np40001,
		Patient:        pt2,
		PatientName:    pt2.Patient_name,
	}
	db.Model(&Admission{}).Create(&ad2)

	ad3 := Admission{
		AdmitTime:      time.Now(),
		RightTreatment: IV0002,
		Room:           pm30001,
		Patient:        pt3,
		PatientName:    pt3.Patient_name,
	}
	db.Model(&Admission{}).Create(&ad3)

	//3.-----------------Treatment---------------------
	//TreatmentRecord
	tr1 := TreatmentRecord{
		RecordDate: time.Now(),
		Treatment:  "Heart Transplant",
		Food_type:  "ควรทานอาหารเหลว งดน้ำตาลในอาหารของผู้ป่วย",
		Med_amount: 3,
		Cost:       50000,

		Equipment: equip2,
		Doctor:    doctor1,
		Med:       med1,
		Admission: ad1,
	}
	db.Model(&TreatmentRecord{}).Create(&tr1)

	tr2 := TreatmentRecord{
		RecordDate: time.Now(),
		Treatment:  "Gastric Lavage",
		Food_type:  "ควรทานอาหารอ่อน เน้นโปรตีนให้ผู้ป่วยเป็นหลัก",
		Med_amount: 3,
		Cost:       50000,

		Equipment: equip1,
		Doctor:    doctor1,
		Med:       med1,
		Admission: ad2,
	}
	db.Model(&TreatmentRecord{}).Create(&tr2)

	tr3 := TreatmentRecord{
		RecordDate: time.Now(),
		Treatment:  "Physical Therapy",
		Food_type:  "ทานอาหารได้ปกติ",
		Med_amount: 4,
		Cost:       30000,

		Equipment: equip2,
		Doctor:    doctor2,
		Med:       med2,
		Admission: ad3,
	}
	db.Model(&TreatmentRecord{}).Create(&tr3)

	//4.-----------------Financial---------------------
	//------------Bill-------------------------
	bill1 := Bill{
		BillDateTime:    time.Now(),
		TreatmentRecord: tr1,
		RightTreatment:  NO0001,
		PaymentMethod:   pm1,
		AmountPaid:      50000,
	}
	db.Model(&Bill{}).Create(&bill1)
	/*bill2 := Bill{
		BillDateTime:    time.Now(),
		TreatmentRecord: tr2,
		RightTreatment:  IV0003,
		PaymentMethod:   pm2,
		AmountPaid:      10000,
	}
	db.Model(&Bill{}).Create(&bill2)*/

	//5.-----------------EquipmentSystem---------------------
	//-------------RequisitionRec-----------
	re1 := RequisitionRecord{
		RecTime:     time.Now(),
		EquipAmount: 1,
		Doctor:      doctor1,
		Equipment:   equip1,
		Admission:   ad1,
		EquipCost:   equip1.Equipment_cost,
	}
	db.Model(&RequisitionRecord{}).Create(&re1)
	/*re2 := RequisitionRecord{
		RecTime:     time.Now(),
		EquipAmount: 2,
		Doctor:      doctor2,
		Equipment:   equip2,
		Admission:   ad2,
		EquipCost:   equip2.Equipment_cost,
	}
	db.Model(&RequisitionRecord{}).Create(&re2)*/

	//6.-----------------FoodAllocate---------------------
	//----------------FoodAllocate----------------------

	// Foodallocate 1
	fa1 := Foodallocate{
		Nutritionist: n1,
		Treatment:    tr1,
		AdmissionID:  tr1.AdmissionID,
		Foodset:      set3,
		Foodtime:     time1,
	}
	db.Model(&Foodallocate{}).Create(&fa1)
	fa2 := Foodallocate{
		Nutritionist: n2,
		Treatment:    tr2,
		AdmissionID:  tr2.AdmissionID,
		Foodset:      set2,
		Foodtime:     time3,
	}
	db.Model(&Foodallocate{}).Create(&fa2)
	fa3 := Foodallocate{
		Nutritionist: n1,
		Treatment:    tr3,
		AdmissionID:  tr3.AdmissionID,
		Foodset:      set2,
		Foodtime:     time2,
	}
	db.Model(&Foodallocate{}).Create(&fa3)

	//7.--------------------MedicationRecord-----------
	//--------------------MedicationRecord------------
	mr1 := MedicationRecord{
		Amount:      3,
		RecordTime:  time.Now(),
		Pharma:      pharma1,
		Med:         med1,
		Treatment:   tr1,
		AdmissionID: tr1.AdmissionID,
	}
	db.Model(&MedicationRecord{}).Create(&mr1)
	/*mr2 := MedicationRecord{
		Amount:      2,
		RecordTime:  time.Now(),
		Pharma:      pharma2,
		Med:         med2,
		Treatment:   tr2,
		AdmissionID: tr2.AdmissionID,
	}
	db.Model(&MedicationRecord{}).Create(&mr2)*/
}
