package api

import (
	"github.com/gin-gonic/gin"
	"learn-crud-hospitals/internal/patient"
	"strconv"
)

// create struct untuk handler pasien
type PatientHandler struct {
	service *patient.PatientService
}

func NewpatientHandler(service *patient.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}

func (h *PatientHandler) GetPatients(c *gin.Context) {
	patients, err := h.service.GetPatients()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, patients)
}

func (h *PatientHandler) GetPatientByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	patient, err := h.service.GetPatientByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, patient)
}

// create handler untuk membuat data pasien
func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patientData patient.Patient //tampung struct patient ke dalam var patientData

	if err := c.ShouldBindJSON(&patientData); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreatePatient(&patientData); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, patientData)
}
