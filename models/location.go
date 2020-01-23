package models

type Location struct {
	LocationId   int    `db:"id" json:"id"`
	LocationName string `db:"name" json:"name"`
	HospitalId   int    `db:"hospital_id" json:"hospital_id"`
}
