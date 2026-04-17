package web

import (
	"go-gin-postgresql-water-monitoring-system-api/internal/usecase/dto"
	"go-gin-postgresql-water-monitoring-system-api/internal/usecase/writeonlyrepository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type input struct {
	OwnerName string `json:"owner" binding:"required,min=3,max=100"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,min=11,max=11"`
	CEP       string `json:"cep" binding:"required,min=8,max=8"`
}

func (i *input) ToDTO() dto.WaterMonitoringInput {
	return dto.WaterMonitoringInput{
		OwnerName: i.OwnerName,
		Email:     i.Email,
		Phone:     i.Phone,
		CEP:       i.CEP,
	} 
}
type output struct {
	ID           int64
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

func ToResponse(dto dto.WaterMonitoringOutput) output {
	return output{
		ID:           dto.ID,
		OwnerName:    dto.OwnerName,
		Email:        dto.Email,
		Phone:        dto.Phone,
		CEP:          dto.CEP,
		PublicPlace:  dto.PublicPlace,
		Neighborhood: dto.Neighborhood,
		State:        dto.State,
		City:         dto.City,
		Region:       dto.Region,
	}
}

type UseCaseRepository interface {
	writeonlyrepository.ICreateUseCase
	writeonlyrepository.IUpdateOwnerUseCase
}

type Handler struct {
	usecase UseCaseRepository
}

func NewHandler(useCaseRepository UseCaseRepository) *Handler {
	return &Handler{
		usecase: useCaseRepository,
	}
}

func (r *Handler) AddNewResidence(ctx *gin.Context) {
	var inp input
	if err := ctx.ShouldBindJSON(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := r.usecase.ExecCreate(inp.ToDTO())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, ToResponse(response))
}

func (r *Handler) UpdateOwner(ctx *gin.Context) {
	var inp input
	if err := ctx.ShouldBindJSON(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	response, err := r.usecase.ExecUpdateOwner(inp.ToDTO(), int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ToResponse(response))
}