package handler

import (
	"capstone-tickets/features/admins"
	"capstone-tickets/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService admins.AdminServiceInterface
}

func New(service admins.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{
		AdminService: service,
	}
}
func (handler *AdminHandler) Register(c echo.Context) error {

	var Register AdminRegister
	errBind := c.Bind(&Register)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}
	var input = AdminRequestToCore(Register)
	err := handler.AdminService.Register(input)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, helpers.Error404+" account not found", nil))
		}
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" password invalid", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", nil))
}
