package handler

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/partners"
	"capstone-tickets/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type PartnerHandler struct {
	PartnerService partners.PartnerServiceInterface
}

func New(service partners.PartnerServiceInterface) *PartnerHandler {
	return &PartnerHandler{
		PartnerService: service,
	}
}

func (handler *PartnerHandler) Login(c echo.Context) error {
	var login PartnerLoginrequest
	errBind := c.Bind(&login)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}
	id, name, token, err := handler.PartnerService.Login(login.Email, login.Password)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, helpers.Error404+" account not found", nil))
		}
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" password invalid", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var data = map[string]any{
		"id":    id,
		"name":  name,
		"token": token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", data))
}

func (handler *PartnerHandler) GetAll(c echo.Context) error {
	var pageParam, itemParam, searchParam string
	_, role := middlewares.ExtractToken(c)

	pageParam = c.QueryParam("page")
	itemParam = c.QueryParam("item")
	searchParam = c.QueryParam("search")

	if role != "Admin" && role != "Superadmin" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	result, next, err := handler.PartnerService.GetAll(pageParam, itemParam, searchParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500+" "+err.Error(), nil))
	}
	var partnerResp = ListPartnerCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.FindAllWebResponse(http.StatusOK, "operation success", partnerResp, next))
}

func (handler *PartnerHandler) Add(c echo.Context) error {
	var partnerRequest PartnerRequest
	var filename string
	file, header, errFile := c.Request().FormFile("profile_picture")
	if errFile != nil {
		if strings.Contains(errFile.Error(), "no such file") {
			filename = helpers.DefaultFile
		} else {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" "+errFile.Error(), nil))
		}
	}

	if filename == "" {
		filename = strings.ReplaceAll(header.Filename, " ", "_")
	}

	errBind := c.Bind(&partnerRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	var partnerCore = PartnerRequestToCore(partnerRequest)
	partnerCore.ProfilePicture = filename

	err := handler.PartnerService.Add(partnerCore, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "operation success", nil))
}

func (handler *PartnerHandler) Get(c echo.Context) error {
	idParam := c.Param("partner_id")
	result, err := handler.PartnerService.Get(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var partnerResponse = PartnerCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", partnerResponse))
}

func (handler *PartnerHandler) Delete(c echo.Context) error {
	id, _ := middlewares.ExtractToken(c)
	idParam := c.Param("partner_id")
	if id != idParam {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	err := handler.PartnerService.Delete(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", nil))
}

func (handler *PartnerHandler) Update(c echo.Context) error {
	id, _ := middlewares.ExtractToken(c)
	idParam := c.Param("partner_id")
	if id != idParam {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}

	var filename string
	file, header, errFile := c.Request().FormFile("profile_picture")
	if errFile != nil {
		if strings.Contains(errFile.Error(), "no such file") {
			filename = helpers.DefaultFile
		} else {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" "+errFile.Error(), nil))
		}
	}

	if filename == "" {
		filename = strings.ReplaceAll(header.Filename, " ", "_")
	}

	var partnerReq PartnerRequest
	errBind := c.Bind(&partnerReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	var partnerCore = PartnerRequestToCore(partnerReq)
	partnerCore.ProfilePicture = filename

	err := handler.PartnerService.Update(idParam, partnerCore, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", nil))
}

func (hendler *PartnerHandler) Test(c echo.Context) error {
	id, role := middlewares.ExtractToken(c)
	data := map[string]any{
		"id":   id,
		"role": role,
	}
	return c.JSON(http.StatusOK, data)
}
