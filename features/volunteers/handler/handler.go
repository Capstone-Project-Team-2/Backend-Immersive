package handler

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/volunteers"
	"capstone-tickets/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

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

	id, name, token, err := h.volunteerService.Login(login.Email, login.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}
	var data = map[string]any{
		"id":    id,
		"name":  name,
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
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	newInput := VolunteerRequestToCore(volunteerReq)

	err := h.volunteerService.Create(newInput)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
		}
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "operation success", nil))
}

func (h *VolunteerHandler) GetAll(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)
	if role != "Partner" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	var qParam volunteers.QueryParam
	page := c.QueryParam("page")
	limitPerPage := c.QueryParam("limitPerPage")

	if limitPerPage == "" {
		qParam.ExistOtherPage = false
	} else {
		qParam.ExistOtherPage = true
		itemsConv, errItem := strconv.Atoi(limitPerPage)
		if errItem != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
		}
		qParam.LimitPerPage = itemsConv
	}
	if page == "" {
		qParam.Page = 1
	} else {
		pageConv, errPage := strconv.Atoi(page)
		if errPage != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
		}
		qParam.Page = pageConv
	}
	searchName := c.QueryParam("search_name")
	qParam.SearchName = searchName
	idParam := c.Param("event_id")
	bol, data, err := h.volunteerService.GetAll(idParam, qParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var volunteerResp = ListVolunteerCoreToResponse(data)
	return c.JSON(http.StatusOK, helpers.FindAllWebResponse(http.StatusOK, "operation success", volunteerResp, bol))
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
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	updatedData := VolunteerRequestToCore(volunteerReq)

	err := h.volunteerService.UpdateById(idParam, updatedData)
	if err != nil {
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
