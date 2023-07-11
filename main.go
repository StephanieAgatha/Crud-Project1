package main

import (
	"github.com/gin-gonic/gin"
	"learn-crud-hospitals/api"
	"learn-crud-hospitals/internal/patient"
	"learn-crud-hospitals/repository"
)

func main() {
	r := gin.Default()

	patientRepo := repository.NewPatientRepository()
	patientService := patient.NewPatientService(patientRepo)
	patientHandler := api.NewpatientHandler(patientService)

	r.GET("/patients", patientHandler.GetPatients)
	r.GET("/patient/:id", patientHandler.GetPatientByID)

	r.Run(":3000")
}
