package main

import (
	hospitalHandler "github.com/NASA12345/google/hospital/delivery"
	hospitalRepository "github.com/NASA12345/google/hospital/repository"
	hospitalUsecase "github.com/NASA12345/google/hospital/usecase"
	locationHandler "github.com/NASA12345/google/location/delivery"
	locationRepository "github.com/NASA12345/google/location/repository"
	locationUsecase "github.com/NASA12345/google/location/usecase"
	patientHandler "github.com/NASA12345/google/patient/delivery"
	patientRepository "github.com/NASA12345/google/patient/repository"
	patientUsecase "github.com/NASA12345/google/patient/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


func main() {
	var db *sqlx.DB
	router := gin.Default()
	var err error
	db = sqlx.MustConnect("postgres", "postgres://meclpohq:h6jPxvrJQkTmTYRZge1f0ZSShyEjItzI@rajje.db.elephantsql.com:5432/meclpohq")
	if err != nil {
		panic(err)
	}
	hospitalRepositor := hospitalRepository.NewhospitalRepository(db)
	locationRepositor := locationRepository.NewLocationRepository(db)
	patientRepositor := patientRepository.NewPatientRepository(db)
	hospitalUsecase := hospitalUsecase.NewHospitalUsecase(hospitalRepositor)
	locationUsecase := locationUsecase.NewLocationUsecase(locationRepositor)
	patientUsecase := patientUsecase.NewPatientUsecase(patientRepositor)
	hospitalGroup := router.Group("/hospitals")
	locationGroup := router.Group("/locations")
	patientGroup := router.Group("/patients")
	hospitalHandler.NewHospitalHandler(hospitalGroup, hospitalUsecase)
	locationHandler.NewLocationHandler(locationGroup, locationUsecase)
	patientHandler.NewPatientHandler(patientGroup, patientUsecase)
	err = router.Run()
}
