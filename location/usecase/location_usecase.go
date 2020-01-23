package usecase

import (
	"github.com/NASA12345/google/location"
	"github.com/NASA12345/google/models"
)


type locationUsecase struct {
	locationRepo    location.Repository
}

func NewLocationUsecase(repository location.Repository) location.Usecase {
	return &locationUsecase{
		locationRepo: repository,
	}
}

func (l locationUsecase) Fetch() ([]models.Location, error){
	return l.locationRepo.Fetch()
}

func (l locationUsecase) GetById(id int64) (models.Location, error) {
	return l.locationRepo.GetById(id)
}

func (l locationUsecase) New(location models.Location) error {
	return l.locationRepo.New(location)
}
