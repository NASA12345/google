package repository

import(
	"github.com/NASA12345/google/models"
	"github.com/NASA12345/google/location"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type LocationRepository struct {
	Db *sqlx.DB
}
/* Create new location repository */
func NewLocationRepository(db *sqlx.DB) location.Repository {
	return &LocationRepository{
		Db: db,
	}
}
/* Get all location data from databse */
func (p *LocationRepository) Fetch() ([]models.Location, error) {
	locations := []models.Location{}
	err := p.Db.Select(&locations, "SELECT * FROM location")
	return locations, err
}
/* Get location with a specific ID */
func (p *LocationRepository) GetById(id int64) (models.Location, error) {
	location := models.Location{}
	err := p.Db.Get(&location, "SELECT * FROM location where id = $1", strconv.FormatInt(id, 10))
	return location, err
}
/* Insert new location in database */
func (p *LocationRepository) New(location models.Location) error {
	_, err := p.Db.NamedExec("INSERT INTO location (name, hospital_id) VALUES (:name, :hospital_id)", location)
	return err
}
