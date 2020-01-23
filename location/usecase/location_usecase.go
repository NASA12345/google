package repository

import(
	"github.com/NASA12345/google/models"
	"github.com/NASA12345/google/location"
)

type LocationUsecase struct {
	Repo location.Repository
}

func NewLocationUsecase(repository location.Repository) location.Usecase {
	return &LocationUsecase{
		Repo: repository,
	}
}

func (p *LocationUsecase) Fetch() ([]models.Location, error) {
	return p.Repo.Fetch()
}

func (p *LocationUsecase) GetById(id int64) (models.Location, error) {
	return p.Repo.GetById(id)
}

func (p *LocationUsecase) New(location models.Location) error {
	return p.Repo.New(location)
}
