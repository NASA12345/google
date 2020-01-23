package patient

import (
	"github.com/NASA12345/google/models"
)

type Usecase interface {
	Fetch() ([]models.Patient, error)
	GetById(id int64) (models.Patient, error)
	New(patient models.Patient) error
}
