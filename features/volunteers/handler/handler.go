package handler

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/volunteers"
	"capstone-tickets/helpers"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var log = helpers.Log()

type VolunteerHandler struct {
	volunteerService volunteers.VolunteerServiceInterface
}

func New(service volunteers.VolunteerServiceInterface) *VolunteerHandler {
	return &VolunteerHandler{
		volunteerService: service,
	}
}

func (h *VolunteerHandler) Login(c echo.Context) error {
	var login LoginRequest
	err := c.Bind(&login)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	id, token, err := h.volunteerService.Login(login.Email, login.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}
	var data = map[string]any{
		"id":    id,
		"token": token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Login success", data))
}

func (h *VolunteerHandler) Create(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)
	if role != "Partner" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	var volunteerReq VolunteerRequest

	errBind := c.Bind(&volunteerReq)
	if errBind != nil {
		log.Error("handler - error on bind request")
		fmt.Println(errBind)
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	newInput := VolunteerRequestToCore(volunteerReq)

	err := h.volunteerService.Create(newInput)
	if err != nil {
		log.Error("handler-internal server error")
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "operation success", nil))

}
func (h *VolunteerHandler) GetAll(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)
	if role != "Partner" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	idParam := c.Param("event_id")
	result, err := h.volunteerService.GetAll(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var partnerResp = ListVolunteerCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.FindAllWebResponse(http.StatusOK, "operation success", partnerResp, false))
}

func (h *VolunteerHandler) GetById(c echo.Context) error {
	id, role := middlewares.ExtractToken(c)
	idParam := c.Param("volunteer_id")
	if id != idParam && role != "Partner" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	result, err := h.volunteerService.GetById(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var volunteerResponse = VolunteerCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", volunteerResponse))
}

func (h *VolunteerHandler) UpdateById(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)
	if role != "Partner" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	idParam := c.Param("volunteer_id")
	var volunteerReq VolunteerRequest
	errBind := c.Bind(&volunteerReq)
	if errBind != nil {
		log.Error("handler - error on bind request")
		fmt.Println(errBind)
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	updatedData := VolunteerRequestToCore(volunteerReq)

	err := h.volunteerService.UpdateById(idParam, updatedData)
	if err != nil {
		log.Error("handler-internal server error")
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", nil))
}

func (h *VolunteerHandler) DeleteById(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)
	if role != "Partner" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	idParam := c.Param("volunteer_id")
	err := h.volunteerService.DeleteById(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", nil))
}
