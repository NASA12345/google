package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/NASA12345/google/location"
	"github.com/NASA12345/google/models"
	"net/http"
	"strconv"
)

type locationHandler struct {
	locationUsecase location.Usecase
}

func (p locationHandler) Fetch(context *gin.Context) {
	location, err := p.locationUsecase.Fetch()
	if err !=nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : err})
	}
	context.JSON(http.StatusOK, location)
}

func (p locationHandler) GetByID(context *gin.Context) {
	tempId, err := strconv.Atoi(context.Param("id"))
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : err})
	}
	id := int64(tempId)
	location, err := p.locationUsecase.GetById(id)
	context.JSON(http.StatusOK, location)

}

func (p locationHandler) NewLocation(context *gin.Context) {
	location := models.Location{}
	err := context.BindJSON(&location)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : err})
	}
	err = p.locationUsecase.New(location)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : err})
	}
	context.JSON(http.StatusOK, gin.H{"message" : "Success!"})
}

func NewLocationHandler(router *gin.RouterGroup, usecase location.Usecase){
	handler := &locationHandler{
		locationUsecase: usecase,
	}
	router.GET("", handler.Fetch)
	router.GET("/:id", handler.GetByID)
	router.POST("", handler.NewLocation)
}
