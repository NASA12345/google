package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)
// Database information inserted
const (
	DbUser     = "meclpohq"
	DbPassword = "h6jPxvrJQkTmTYRZge1f0ZSShyEjItzI"
	DbName     = "meclpohq host=rajje.db.elephantsql.com"
	DbHost     = "rajje.db.elephantsql.com"
)

var router *gin.Engine
var db *sqlx.DB
// Structure of database is declared
type patient struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Illness string `json:"illness" db:"illness"`
	BirthDate string `json:"birthDate" db:"birth_date"`
	LocationId int `json:"locationId" db:"location_id"`
}

type location struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	HospitalId string `json:"hospitalId" db:"hospital_id"`
}

type hospital struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	MaxPatientAmount int `json:"maxPatientCount" db:"max_patient_amount"`
}

// function declaring abouter router
func main() {
	router = gin.Default()
	initDatabase()
	initAPI()
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln("Could not run the app")
	}
	defer db.Close()
}

func initDatabase() {
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", DbUser, DbPassword, DbName, DbHost)
	db= sqlx.MustConnect("postgres", psqlInfo)
}
// Provide details of all patients
func GetAllPatients(c *gin.Context) {
	var patients []patient
	err := db.Select(&patients, "select * from patient")
	if err != nil {
		log.Panicln(err)
	}
	c.JSON(http.StatusOK, patients)
}


// Provide deatils of all locations
func GetAllLocations(c *gin.Context) {
	var locations []location
	err := db.Select(&locations, "select * from location")
	if err != nil {
		log.Panicln(err)
	}
	c.JSON(http.StatusOK, locations)
}

// Provoide details of all hospitals
func GetAllHospitals(c *gin.Context) {
	var hospitals []hospital
	err := db.Select(&hospitals, "select * from hospital")
	if err != nil {
		log.Panicln(err)
	}
	c.JSON(http.StatusOK, hospitals)
}

// This function will initialize all the routes
func initAPI() {
	router.GET("/patient/all", GetAllPatients)
	router.GET("/location/all", GetAllLocations)
	router.GET("/hospital/all", GetAllHospitals)
}
