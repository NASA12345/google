package location

import (
	    "github.com/NASA12345/google/models"
)
type Usecase interface {
	Fetch() ([]models.Location, error)
	GetById(id int64) (models.Location, error)
	New(hospital models.Location) error
}
