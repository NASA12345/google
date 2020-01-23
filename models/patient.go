package models

type Patient struct {
	PatientId   int    `db:"id" json:"id"`
	PatientName string `db:"name" json:"name"`
	Illness     string `db:"illness" json:"illness"`
	BirthDate   string `db:"birth_date" json:"birth_date"`
	LocationId  int    `db:"location_id" json:"location_id"`
}
