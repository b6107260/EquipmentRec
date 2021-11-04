package main

import (
	actor "github.com/ProjectG10/controller/actor"
	admission "github.com/ProjectG10/controller/admission"
	admission_system "github.com/ProjectG10/controller/admission_system"
	disease "github.com/ProjectG10/controller/disease"
	equipment "github.com/ProjectG10/controller/equipment"
	equipment_system "github.com/ProjectG10/controller/equipment_system"
	food_system "github.com/ProjectG10/controller/food_system"
	medication_record_system "github.com/ProjectG10/controller/medication_record_system"
	medicine "github.com/ProjectG10/controller/medicine"
	patient "github.com/ProjectG10/controller/patient"
	payment_system "github.com/ProjectG10/controller/payment_system"
	right_trearment "github.com/ProjectG10/controller/right_treatment"
	treatment_record "github.com/ProjectG10/controller/treatment_record"

	"github.com/ProjectG10/entity"
	"github.com/ProjectG10/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	/*gin.SetMode(gin.ReleaseMode)*/
	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			//routes admission--------major
			protected.GET("/route/GetAdmission/:id", admission.GetAdmission)
			protected.GET("/route/ListAdmission", admission.ListAdmission)
			protected.POST("/route/CreatAdmission", admission.CreateAdmission)
			//routes admission_system/roomtypes
			protected.GET("/route/GetRoomType/:id", admission_system.GetRoomtype)
			protected.GET("/route/ListRoomType", admission_system.ListRoomtypes)
			//routes admission_system/room
			protected.GET("/route/GetRoom/:id", admission_system.GetRoom)
			protected.GET("/route/ListRoom", admission_system.ListRoom)
			//route equipment
			protected.GET("/route/GetEquipment/:id", equipment.GetEquipment)
			protected.GET("/route/ListEquipment", equipment.ListEquipment)
			//route equipment system/requisition--------major
			protected.GET("/route/GetRequisition/:id", equipment_system.GetRequisitionRecord)
			protected.GET("/route/ListRequisition", equipment_system.ListRequisitionRecord)
			protected.POST("/route/CreatRequisition", equipment_system.CreateRequisitionRecord)
			//route food_system/food allocate--------major
			protected.GET("/route/GetFoodAllocate/:id", food_system.GetFoodallocate)
			protected.GET("/route/ListFoodAllocate", food_system.ListFoodallocates)
			protected.POST("/route/CreatFoodAllocate", food_system.CreateFoodallocate)
			//route food_system/food set
			protected.GET("/route/GetFoodSet/:id", food_system.GetFoodSet)
			protected.GET("/route/ListFoodSet", food_system.ListFoodSets)
			//route food_system/food time
			protected.GET("/route/GetFoodTime/:id", food_system.GetFoodTime)
			protected.GET("/route/ListFoodTime", food_system.ListFoodTimes)
			//route medication_record_system-----------major
			protected.GET("/route/GetMedRec/:id", medication_record_system.GetMedicationRacord)
			protected.GET("/route/ListMedRec", medication_record_system.ListMedicationRacord)
			protected.POST("/route/CreatMedRec", medication_record_system.CreateMedicationRecord)
			//route medicine
			protected.GET("/route/GetMedicine/:id", medicine.GetMedicine)
			protected.GET("/route/ListMedicine", medicine.ListMedicine)
			//route patient--------------major
			protected.GET("/route/GetPatient/:id", patient.GetPatient)
			protected.GET("/route/ListPatient", patient.ListPatients)
			protected.POST("/route/CreatPatient", patient.CreatePatient)
			//route disease
			protected.GET("/route/GetDisease/:id", disease.GetDisease)
			protected.GET("/route/ListDisease", disease.ListDiseases)
			//route payment_system/bill-------------major
			protected.GET("/route/GetBill/:id", payment_system.GetBill)
			protected.GET("/route/ListBill", payment_system.ListBill)
			protected.POST("/route/CreatBill", payment_system.CreateBill)
			//route payment_system/payment
			protected.GET("/route/GetPayment/:id", payment_system.GetPaymentMethod)
			protected.GET("/route/ListPayment", payment_system.ListPaymentMethod)
			//route right_treatment
			protected.GET("/route/GetRightTreatment/:id", right_trearment.GetRightTreatment)
			protected.GET("/route/ListRightTreatment", right_trearment.ListRightTreatment)
			//route treatment_record_system------------major
			protected.GET("/route/GetTreatmentRec/:id", treatment_record.GetTreatmentRecord)
			protected.GET("/route/ListTreatmentRec", treatment_record.ListTreatmentRecord)
			protected.POST("/route/CreatTreatmentRec", treatment_record.CreateTreatmentRecord)
			//route role
			// Pharmacist Routes
			protected.GET("/route/ListPharmacist", actor.ListPharmacists)
			protected.GET("/route/GetPharmacist/:id", actor.GetPharmacist)
			protected.POST("/route/CreatePharmacist", actor.CreatePharmacist)
			// Nurse Routes
			protected.GET("/route/ListNurse", actor.ListNurses)
			protected.GET("/route/GetNurse/:id", actor.GetNurse)
			protected.POST("/route/CreateNurse", actor.CreateNurse)
			// Doctor Routes
			protected.GET("/route/ListDoctor", actor.ListDoctors)
			protected.GET("/route/GetDoctor/:id", actor.GetDoctor)
			protected.POST("/route/CreateDoctor", actor.CreateDoctor)
			// Nutritionist Routes
			protected.GET("/route/ListNutritionist", actor.ListNutritionists)
			protected.GET("/route/GetNutritionist/:id", actor.GetNutritionist)
			protected.POST("/route/CreateNutritionist", actor.CreateNutritionist)
			// FinancialOfficer Routes
			protected.GET("/route/ListFinancialOfficer", actor.ListFinancialOfficer)
			protected.GET("/route/GetFinancialOfficer/:id", actor.GetFinancialOfficer)
			protected.POST("/route/CreateFinancialOfficer", actor.CreateFinancialOfficer)

		}
	}
	// Actor Routes
	r.POST("/pharmacists/create", actor.CreatePharmacist)
	r.POST("/nurses/create", actor.CreatePharmacist)
	r.POST("/financialofficers/create", actor.CreatePharmacist)
	r.POST("/nutritionists/create", actor.CreatePharmacist)
	r.POST("/doctors/create", actor.CreatePharmacist)

	// Login/Actor
	r.POST("/login/pharmacist", actor.LoginPharmacist)
	r.POST("/login/nurse", actor.LoginNurse)
	r.POST("/login/financialofficer", actor.LoginFinancialOfficer)
	r.POST("/login/nutritionist", actor.LoginNurse)
	r.POST("/login/doctor", actor.LoginDoctor)

	// Run the server
	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	}

}
