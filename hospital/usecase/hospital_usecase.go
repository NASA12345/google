package repository

import(
	"github.com/NASA12345/google/models"
	"github.com/NASA12345/google/hospital"
)

type HospitalUsecase struct {
	Repo hospital.Repository
}

func NewHospitalUsecase(repository hospital.Repository) hospital.Usecase {
	return &HospitalUsecase{
		Repo: repository,
	}
}

func (p *HospitalUsecase) Fetch() ([]models.Hospital, error) {
	return p.Repo.Fetch()
}

func (p *HospitalUsecase) GetById(id int64) (models.Hospital, error) {
	return p.Repo.GetById(id)
}

func (p *HospitalUsecase) New(hospital models.Hospital) error {
	return p.Repo.New(hospital)
}
