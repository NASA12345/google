package models

type Hospital struct {
	HospitalId       int    `db:"id" json:"id"`
	HospitalName     string `db:"name" json:"name"`
	MaxPatientAmount int    `db:"max_patient_amount" json:"max_patient_amount"`
}
