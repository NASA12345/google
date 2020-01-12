package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
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
var db *sql.DB
// Structure of database is declared
type patient struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Illness    string `json:"illness"`
	BirthDate  string `json:"birth_date"`
	LocationId int    `json:"location_id"`
}

type location struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	HospitalId int    `json:"hospital_id"`
}

type hospital struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	HospitalId int    `json:"max_patient_amount"`
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
	var err error

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to the database")
}
// Provide details of all patients
func GetAllPatients(c *gin.Context) {

	rows, err := db.Query("SELECT * FROM patient")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var patients []patient
	for rows.Next() {
		row := patient{}
		err = rows.Scan(&row.ID, &row.Name, &row.Illness, &row.BirthDate, &row.LocationId)

		if err != nil {
			panic(err)
		}
		patients = append(patients, row)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{"patients": patients})
}

// Provide deatils of all locations
func GetAllLocations(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM location")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var locations []location
	for rows.Next() {
		row := location{}
		err = rows.Scan(&row.ID, &row.Name, &row.HospitalId)

		if err != nil {
			panic(err)
		}
		locations = append(locations, row)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{"locations": locations})
}

// Provoide details of all hospitals
func GetAllHospitals(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM hospital")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var hospitals []hospital
	for rows.Next() {
		row := hospital{}
		err = rows.Scan(&row.ID, &row.Name, &row.HospitalId)

		if err != nil {
			panic(err)
		}
		hospitals = append(hospitals, row)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{"hospitals": hospitals})
}

// This function will initialize all the routes
func initAPI() {
	router.GET("/patient/all", GetAllPatients)
	router.GET("/location/all", GetAllLocations)
	router.GET("/hospital/all", GetAllHospitals)
}
