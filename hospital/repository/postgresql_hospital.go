package repository

import(
	"github.com/NASA12345/google/models"
	"github.com/NASA12345/google/hospital"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type HospitalRepository struct {
	Db *sqlx.DB
}
/* create new hospital repository*/
func NewHospitalRepository(db *sqlx.DB) hospital.Repository {
	return &HospitalRepository{
		Db: db,
	}
}
/* Get all hospital data from database*/
func (p *HospitalRepository) Fetch() ([]models.Hospital, error) {
	hospitals := []models.Hospital{}
	err := p.Db.Select(&hospitals, "SELECT * FROM hospital")
	return hospitals, err
}
/* Get hospital with a specific ID */
func (p *HospitalRepository) GetById(id int64) (models.Hospital, error) {
	hospital := models.Hospital{}
	err := p.Db.Get(&hospital, "SELECT * FROM hospital where id = $1", strconv.FormatInt(id, 10))
	return hospital, err
}
/* Insert new hospital in database */
func (p *HospitalRepository) New(hospital models.Hospital) error {
	_, err := p.Db.NamedExec("INSERT INTO hospital (name, max_patient_amount) VALUES (:name, :max_patient_amount)", hospital)
	return err
}
