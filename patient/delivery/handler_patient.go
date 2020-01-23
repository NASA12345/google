package delivery

import (
	"github.com/NASA12345/google/models"
	"github.com/NASA12345/google/patient"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type patientHandler struct {
	patientUsecase patient.Usecase
}

func (p patientHandler) Fetch(context *gin.Context) {
	patient, err := p.patientUsecase.Fetch()
	if err !=nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : err})
	}
	context.JSON(http.StatusOK, patient)
}

func (p patientHandler) GetByID(context *gin.Context) {
	tempId, err := strconv.Atoi(context.Param("id"))
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : err})
	}
	id := int64(tempId)
	patient, err := p.patientUsecase.GetById(id)
	context.JSON(http.StatusOK, patient)

}

func (p patientHandler) NewPatient(context *gin.Context) {
	patient := models.Patient{}
	err := context.Bind(&patient)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : err})
	}
	err = p.patientUsecase.New(patient)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : err})
	}
	context.JSON(http.StatusOK, gin.H{"message" : "Success!"})
}

func NewPatientHandler(router *gin.RouterGroup, usecase patient.Usecase){
	handler := &patientHandler{
		patientUsecase: usecase,
	}
	router.GET("", handler.Fetch)
	router.GET("/:id", handler.GetByID)
	router.POST("", handler.NewPatient)
}
