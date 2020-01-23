package usecase

import (
	"github.com/NASA12345/google/models"
	"github.com/NASA12345/google/patient"
)


type patientUsecase struct {
	patientRepo patient.Repository
}

func NewPatientUsecase(repository patient.Repository) patient.Usecase {
	return &patientUsecase{
		patientRepo: repository,
	}
}

func (p patientUsecase) Fetch() ([]models.Patient, error){
	return p.patientRepo.Fetch()
}

func (p patientUsecase) GetById(id int64) (models.Patient, error) {
	return p.patientRepo.GetById(id)
}

func (p patientUsecase) New(location models.Patient) error {
	return p.patientRepo.New(location)
}

