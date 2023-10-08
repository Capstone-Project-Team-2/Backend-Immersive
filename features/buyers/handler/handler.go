package handler

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/buyers"
	"capstone-tickets/helpers"
	"net/http"
	"strconv"
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
	var login LoginRequest
	err := c.Bind(&login)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	id, name, token, err := h.buyerService.Login(login.Email, login.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+"Input tidak valid, harap isi email dan password sesuai ketentuan"+err.Error(), nil))
		} else if strings.Contains(err.Error(), "invalid email") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+"Email yang anda berikan tidak terdaftar"+err.Error(), nil))
		} else {
			return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+"password yang anda berikan tidak valid"+err.Error(), nil))
		}
	}
	var data = map[string]any{
		"id":    id,
		"name":  name,
		"token": token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Login success", data))
}

func (h *BuyerHandler) Create(c echo.Context) error {
	var buyerReq BuyerRequest
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

	errBind := c.Bind(&buyerReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+""+errBind.Error(), nil))
	}

	newInput := BuyerRequestToCore(buyerReq)
	newInput.ProfilePicture = filename

	err := h.buyerService.Create(newInput, file)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500+" "+err.Error(), nil))
		}
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "operation success", nil))
}
func (h *BuyerHandler) GetAll(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)
	if role != "Admin" && role != "Superadmin" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	var qParam buyers.QueryParam
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
	bol, data, err := h.buyerService.GetAll(qParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var buyerResp = ListBuyerCoreToResponse(data)
	return c.JSON(http.StatusOK, helpers.FindAllWebResponse(http.StatusOK, "operation success", buyerResp, bol))
}

func (h *BuyerHandler) GetById(c echo.Context) error {
	id, _ := middlewares.ExtractToken(c)
	idParam := c.Param("buyer_id")
	if id != idParam {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	result, err := h.buyerService.GetById(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var buyerResponse = BuyerCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", buyerResponse))
}

func (h *BuyerHandler) UpdateById(c echo.Context) error {
	id, _ := middlewares.ExtractToken(c)
	idParam := c.Param("buyer_id")
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

	var buyerReq BuyerRequest
	errBind := c.Bind(&buyerReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" "+errBind.Error(), nil))
	}

	updatedData := BuyerRequestToCore(buyerReq)
	updatedData.ProfilePicture = filename

	err := h.buyerService.UpdateById(idParam, updatedData, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", nil))
}

func (h *BuyerHandler) DeleteById(c echo.Context) error {
	id, _ := middlewares.ExtractToken(c)
	idParam := c.Param("buyer_id")
	if id != idParam {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	err := h.buyerService.DeleteById(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", nil))
}
