package main

import (
	"github.com/b6107260/sa-64-equipment/controller"

	"github.com/b6107260/sa-64-equipment/entity"
	"github.com/b6107260/sa-64-equipment/middlewares"

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
			// Doctor Routes
			protected.GET("/doctors", controller.ListDoctors)
			protected.GET("/doctor/:id", controller.GetDoctor)
			protected.POST("/doctors", controller.CreateDoctor)
			//protected.PATCH("/doctors", controller.UpdateDoctor)
			//protected.DELETE("/doctors/:id", controller.DeleteDoctor)

			// Equipment Routes
			protected.GET("/equipments", controller.ListEquipment)
			protected.GET("/equipment/:id", controller.GetEquipment)
			protected.POST("/equipments", controller.CreateEquipment)
			//protected.PATCH("/equipments", controller.UpdateEquipment)
			//protected.DELETE("/equipments/:id", controller.DeleteEquipment)

			// Admission Routes
			protected.GET("/admissions", controller.ListAdmission)
			protected.GET("/admission/:id", controller.GetAdmission)
			//protected.GET("/admission/watched/user/:id", controller.GetPlaylistWatchedByUser)
			protected.POST("/admissions", controller.CreateAdmission)
			//protected.PATCH("/admissions", controller.UpdateAdmission)
			//protected.DELETE("/admissions/:id", controller.DeleteAdmission)

			// ReqRecord Routes
			protected.GET("/requisition_records", controller.ListRequisitionRecord)
			protected.GET("/requisition_record/:id", controller.GetRequisitionRecord)
			protected.POST("/requisition_records", controller.CreateRequisitionRecord)
			//protected.PATCH("/requisition_records", controller.UpdateRequisitionRacord)
			//protected.DELETE("/requisition_records/:id", controller.DeleteRequisitionRacord)

		}
	}
	// Doctor Routes
	r.POST("/doctors/create", controller.CreateDoctor)

	// Authentication Routes
	r.POST("/login", controller.Login)

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