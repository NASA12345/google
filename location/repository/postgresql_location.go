package repository

import (
	"github.com/NASA12345/google/models"
	"github.com/jmoiron/sqlx"
)

type locationRepository struct {
	db *sqlx.DB
}

/* Create new location repository */
func NewLocationRepository(db *sqlx.DB) *locationRepository {
	return &locationRepository{
		db: db,
	}
}


/* Get all location data from databse */
func (m *locationRepository) Fetch() ([]models.Location, error) {
	var location []models.Location
	error := m.db.Select(&location, "SELECT * from location")
	if error != nil {
		panic(error)
	}
	return location, error
}


/* Get location with a specific ID */
func (m *locationRepository) GetById(id int64) (models.Location, error) {
	var location models.Location
	error := m.db.Get(&location,"SELECT * from location WHERE id=$1", id)
	return location, error
}


/* Insert new location in database */
func (m *locationRepository) New(location models.Location) error {
	_, err := m.db.Exec("INSERT into location VALUES($1, $2, $3)", location.LocationId, location.LocationName, location.HospitalId)
	return err
}


