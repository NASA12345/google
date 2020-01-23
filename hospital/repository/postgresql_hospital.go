package repository

import (
	"github.com/NASA12345/google/models"
	"github.com/jmoiron/sqlx"
)

type hospitalRepository struct {
	db *sqlx.DB
}

/* create new hospital repository*/
func NewhospitalRepository(db *sqlx.DB) *hospitalRepository {
	return &hospitalRepository{
		db: db,
	}
}
/* Get all hospital data from database*/
func (m *hospitalRepository) Fetch() ([]models.Hospital, error) {
	var hospital []models.Hospital
	error := m.db.Select(&hospital, "SELECT * from hospital")
	return hospital, error
}

/* Get hospital with a specific ID */
func (m *hospitalRepository) GetById(id int64) (models.Hospital, error) {
	var hospital models.Hospital
	error := m.db.Get(&hospital,"SELECT * from hospital WHERE id=$1", id)
	return hospital, error
}
/* Insert new hospital in database */
func (m *hospitalRepository) New(hospital models.Hospital) error {
	_, err := m.db.Exec("INSERT into hospital VALUES($1, $2, $3)", hospital.HospitalId, hospital.HospitalName, hospital.MaxPatientAmount)
	return err
}

