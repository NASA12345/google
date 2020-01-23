package location

import (
	"github.com/NASA12345/google/models"
)

type Repository interface {
	Fetch() ([]models.Location, error)
	GetById(id int64) (models.Location, error)
	New(location models.Location) error
}
