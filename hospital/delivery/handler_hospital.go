package delivery

import (
	"github.com/NASA12345/google/hospital"
	"github.com/NASA12345/google/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HospitalHandler struct {
	hospitalUsecase hospital.Usecase
}

func (h HospitalHandler) Fetch(context *gin.Context) {
	hospital, err := h.hospitalUsecase.Fetch()
	if err !=nil{
		panic(err)
	}
	context.JSON(http.StatusOK, hospital)
}

func (h HospitalHandler) GetByID(context *gin.Context) {
	tempId, err := strconv.Atoi(context.Param("id"))
	if err != nil{
		panic(err)
	}
	id := int64(tempId)
	hospital, err := h.hospitalUsecase.GetById(id)
	context.JSON(http.StatusOK, hospital)

}

func (h HospitalHandler) NewHospital(context *gin.Context) {
	hospital := models.Hospital{}
	err := context.Bind(&hospital)
	if err != nil{
		context.JSON(http.StatusOK, hospital)
	}
	err = h.hospitalUsecase.New(hospital)
	if err != nil{
		panic(err)
	}
	context.JSON(http.StatusOK, hospital)
}

func NewHospitalHandler(router *gin.RouterGroup, usecase hospital.Usecase){
	handler := &HospitalHandler{
		hospitalUsecase: usecase,
	}
	router.GET("", handler.Fetch)
	router.GET("/:id", handler.GetByID)
	router.POST("", handler.NewHospital )
}
