package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"

	patientR "github.com/NASA12345/google/patient/repository"
	patientU "github.com/NASA12345/google/patient/usecase"
	patientH "github.com/NASA12345/google/patient/delivery"

	locationR "github.com/NASA12345/google/location/repository"
	locationU "github.com/NASA12345/google/location/usecase"
	locationH "github.com/NASA12345/google/location/delivery"

	hospitalR "github.com/NASA12345/google/hospital/repository"
	hospitalU "github.com/NASA12345/google/hospital/usecase"
	hospitalH "github.com/NASA12345/google/hospital/delivery"
)

func main() {
	r := gin.Default()
	url := "postgres://meclpohq:h6jPxvrJQkTmTYRZge1f0ZSShyEjItzI@rajje.db.elephantsql.com:5432/meclpohq"
	db := sqlx.MustConnect("postgres", url)

	patientRepository := patientR.NewPatientRepository(db)
	patientUsecase := patientU.NewPatientUsecase(patientRepository)
	patientHandler := r.Group("/patients/")
	patientH.NewPatientHandler(patientHandler, patientUsecase)

	hospitalRepository := hospitalR.NewHospitalRepository(db)
	hospitalUsecase := hospitalU.NewHospitalUsecase(hospitalRepository)
	hospitalHandler := r.Group("/hospitals/")
	hospitalH.NewHospitalHandler(hospitalHandler, hospitalUsecase)

	locationRepository := locationR.NewLocationRepository(db)
	locationUsecase := locationU.NewLocationUsecase(locationRepository)
	locationHandler := r.Group("/locations/")
	locationH.NewLocationHandler(locationHandler, locationUsecase)

	err := r.Run()
	if err != nil {
		panic(err)
	}

}
