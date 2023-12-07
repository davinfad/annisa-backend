package handler

import (
	"annisa-salon/auth"
	"annisa-salon/helper"
	"annisa-salon/input"
	"annisa-salon/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type treatmentsHandler struct {
	treatmentService service.ServiceTreatments
	authService      auth.UserAuthService
}

func NewTreatmentsHandler(treatmentService service.ServiceTreatments, authService auth.UserAuthService) *treatmentsHandler {
	return &treatmentsHandler{treatmentService, authService}
}

func (h *treatmentsHandler) CreateTreatments (c *gin.Context) {
	var inputTreatments input.InputTreatments

	err := c.ShouldBindJSON(&inputTreatments)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	createTreatment, err := h.treatmentService.CreateTreatment(inputTreatments)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, createTreatment)
	c.JSON(http.StatusOK, response)
}

func (h *treatmentsHandler) UpdatedTreatment(c *gin.Context) {
	slug := c.Param("slug")

	var inputTreatments input.InputTreatments

	err := c.ShouldBindJSON(&inputTreatments)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	UpdatedTreatment, err := h.treatmentService.UpdateTreatment(slug, inputTreatments)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, UpdatedTreatment)
	c.JSON(http.StatusOK, response)
}

func (h *treatmentsHandler) GetAllTreatments (c *gin.Context){
	// slug := c.Param("slug")
	// finalSlug = c.Param("finalSlug")
	
	Blog, err := h.treatmentService.FindAllTreatment()

	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
        c.JSON(http.StatusBadRequest, response)
        return
	}
	
	response := helper.APIresponse(http.StatusOK, Blog)
	c.JSON(http.StatusOK, response)
}

func (h *treatmentsHandler) GetOneTreatment (c *gin.Context) {
	slug := c.Param("slug")

	Blog, err := h.treatmentService.FindTreatmentBySlug(slug)

	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
        c.JSON(http.StatusBadRequest, response)
        return
	}
	
	response := helper.APIresponse(http.StatusOK, Blog)
	c.JSON(http.StatusOK, response)
}

func (h *treatmentsHandler) DeleteTreatment (c *gin.Context) {
	slug := c.Param("slug")

	err := h.treatmentService.DeleteTreatment(slug)

	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
        c.JSON(http.StatusBadRequest, response)
        return
	}
	
	response := helper.APIresponse(http.StatusOK, "blog has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}