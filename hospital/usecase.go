package hospital

import (
	"github.com/NASA12345/google/models"
)

type Usecase interface {
	Fetch() ([]models.Hospital, error)
	GetById(id int64) (models.Hospital, error)
	New(hospital models.Hospital) error
}
