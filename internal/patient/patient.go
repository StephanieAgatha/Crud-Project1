package patient

import "learn-crud-hospitals/repository"

type Patient struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
}

type PatientRepository interface {
	GetAll() ([]Patient, error)
	GetByID(id int) (*Patient, error)
	Create(patient *Patient) error
}

type PatientService struct {
	repo PatientRepository
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func (s *PatientService) GetPatients() ([]Patient, error) {
	return s.repo.GetAll()
}

func (s *PatientService) GetPatientByID(id int) (*Patient, error) {
	return s.repo.GetByID(id)
}

func (s *PatientService) CreatePatient(patient *Patient) error {
	return s.repo.Create(patient)
}
