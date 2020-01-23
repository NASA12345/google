package usecase

import (
	"github.com/NASA12345/google/hospital"
	"github.com/NASA12345/google/models"
)


type hospitalUsecase struct {
	hospitalRepo    hospital.Repository
}

func NewHospitalUsecase(repository hospital.Repository) hospital.Usecase {
	return &hospitalUsecase{
		hospitalRepo: repository,
	}
}

func (h hospitalUsecase) Fetch() ([]models.Hospital, error) {
	return h.hospitalRepo.Fetch()
}

func (h hospitalUsecase) GetById(id int64) (models.Hospital, error) {
	return h.hospitalRepo.GetById(id)
}

func (h hospitalUsecase) New(hospital models.Hospital) error {
	return h.hospitalRepo.New(hospital)
}
