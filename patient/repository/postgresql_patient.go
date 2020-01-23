package repository

import (
	"github.com/NASA12345/google/models"
	"github.com/jmoiron/sqlx"
)

type patientRepository struct {
	db *sqlx.DB
}

/* Create new patient repository */
func NewPatientRepository(db *sqlx.DB) *patientRepository {
	return &patientRepository{
		db: db,
	}
}
/* Create all patient data from database */
func (m *patientRepository) Fetch() ([]models.Patient, error) {
	var patient []models.Patient
	error := m.db.Select(&patient, "SELECT * from patient")
	if error != nil {
		panic(error)
	}
	return patient, error
}

/* Get patient with a specific ID */
func (m *patientRepository) GetById(id int64) (models.Patient, error) {
	var patient models.Patient
	error := m.db.Get(&patient,"SELECT * from patient WHERE id=$1", id)
	return patient, error
}

/* Insert new patient in database */
func (m *patientRepository) New(patient models.Patient) error {
	_, err := m.db.Exec("INSERT into patient VALUES($1, $2, $3, $4, $5)", patient.PatientId, patient.PatientName, patient.Illness, patient.BirthDate, patient.LocationId)
	return err
}
