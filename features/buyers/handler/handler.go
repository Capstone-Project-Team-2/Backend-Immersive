package handler

import (
	"capstone-tickets/features/buyers"
	"capstone-tickets/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type BuyerHandler struct {
	buyerService buyers.BuyerServiceInterface
}

func New(service buyers.BuyerServiceInterface) *BuyerHandler {
	return &BuyerHandler{
		buyerService: service,
	}
}

func (h *BuyerHandler) Login(c echo.Context) error {
	login := new(LoginRequest)
	err := c.Bind(&login)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	id, token, err := h.buyerService.Login(login.Email, login.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}
	var data = map[string]any{
		"id":    id,
		"token": token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", data))
}

func (h *BuyerHandler) Register(c echo.Context) error {
	NewBuyer := new(BuyerRequest)
	var linkFile string
	file, header, errFile := c.Request().FormFile("avatar")
	if errFile != nil {
		if strings.Contains(errFile.Error(), "no such file") {
			linkFile = helpers.DefaultFile
		} else {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" "+errFile.Error(), nil))
		}
	}

	if linkFile == "" {
		linkFile = strings.ReplaceAll(header.Filename, " ", "_")
	}

	err := c.Bind(&NewBuyer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	newInput := BuyerRequestToCore(*NewBuyer)
	newInput.Avatar = linkFile

	err = h.buyerService.Register(newInput, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "operation success", nil))

}
