package web

import (
	"go-gin-postgresql-water-monitoring-system-api/internal/usecase/dto"
	"go-gin-postgresql-water-monitoring-system-api/internal/usecase/writeonlyrepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WaterMonitoringInputDTO struct {
	OwnerName string `json:"owner" binding:"required,min=3,max=100"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,min=11,max=11"`
	CEP       string `json:"cep" binding:"required,min=8,max=8"`
}

type WaterMonitoringOutputDTO struct {
	ID           int
	OwnerName    string
	Email        string
	Phone        string
	CEP          string
	PublicPlace  string
	Neighborhood string
	State        string
	City         string
	Region       string
}

type UseCaseRepository interface {
	writeonlyrepository.ICreateUseCase
}

type Handler struct {
	usecase UseCaseRepository
}

func NewHandler(useCaseRepository UseCaseRepository) *Handler {
	return &Handler{
		usecase: useCaseRepository,
	}
}

func (r *Handler) AddNewAddress(ctx *gin.Context) {
	var input WaterMonitoringInputDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ucDTO := dto.WaterMonitoringInput{
		OwnerName: input.OwnerName,
		Email:     input.Email,
		Phone:     input.Phone,
		CEP:       input.CEP,
	}

	response, err := r.usecase.Execute(ucDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responseDTO := WaterMonitoringOutputDTO{
		ID:           response.ID,
		OwnerName:    response.OwnerName,
		Email:        response.Email,
		Phone:        response.Phone,
		CEP:          response.CEP,
		PublicPlace:  response.PublicPlace,
		Neighborhood: response.Neighborhood,
		State:        response.State,
		City:         response.City,
		Region:       response.Region,
	}

	ctx.JSON(http.StatusOK, responseDTO)
}
