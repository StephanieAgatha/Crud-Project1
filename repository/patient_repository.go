package repository

import (
	"errors"
	"learn-crud-hospitals/internal/patient"
)

type PatientRepository struct {
	patients []patient.Patient //Field ini akan digunakan untuk menampung data pasien.
}

func NewPatientRepository() *PatientRepository {
	initialPatients := []patient.Patient{
		{
			ID:      1,
			Name:    "Je",
			Address: "SG",
			Age:     20,
		},
		{
			ID:      2,
			Name:    "Ka",
			Address: "SG",
			Age:     19,
		},
	}
	return &PatientRepository{patients: initialPatients}
	//mengembalikan pointer ke instance baru dari PatientRepository yang diinisialisasi dengan slice initialPatients yang diberikan pada field patients.
}

// (r *PatientRepository) adalah metode yang terkait dengan tipe PatientRepository.
// r adalah receiver yang mengacu pada instance PatientRepository yang digunakan saat memanggil metode ini.
// func untuk GetAll pasien
func (r *PatientRepository) GetAll() ([]patient.Patient, error) {
	return r.patients, nil
}

// func untuk Get pasien by id,mereturn pointer ke struct patient
func (r *PatientRepository) GetByID(id int) (*patient.Patient, error) {
	//lakukan for range untuk mengakses struct patients
	for _, patient := range r.patients {
		if patient.ID == id {
			return &patient, nil
		}
	}
	//jika pasien tidak ditemukan
	return nil, errors.New("Patient Not Found")
}

// func untuk menambahkan data pasien di dalam patientRepository
// patient *patient.Patient adalah parameter yang diterima oleh metode Create.
// Parameter ini adalah pointer ke objek patient.Patient yang akan dibuat dan disimpan di dalam repository.
func (r *PatientRepository) Create(patient *patient.Patient) error {
	//mengatur nilai ID pada objek patient dengan menggunakan (length) slice r.patients ditambah 1.
	//Ini memberikan ID unik untuk objek pasien baru.
	patient.ID = len(r.patients) + 1
	r.patients = append(r.patients, *patient)
	return nil
}
