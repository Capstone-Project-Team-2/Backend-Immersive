package handler

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/transactions"
	"capstone-tickets/helpers"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var log = helpers.Log()

type TransactionHandler struct {
	transactionService transactions.TransactionServiceInterface
}

func New(service transactions.TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{
		transactionService: service,
	}
}

func (h *TransactionHandler) Create(c echo.Context) error {
	buyer_id, _ := middlewares.ExtractToken(c)
	var transactionReq TransactionRequest

	errBind := c.Bind(&transactionReq)
	if errBind != nil {
		log.Error("handler - error on bind request")
		fmt.Println(errBind)
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	newInput := TransactionRequestToCore(transactionReq)

	err := h.transactionService.Create(newInput, buyer_id)
	if err != nil {
		log.Error("handler-internal server error")
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "operation success", nil))
}

func (h *TransactionHandler) GetById(c echo.Context) error {
	idParam := c.Param("transaction_id")
	result, err := h.transactionService.Get(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var transactionResponse = TransactionCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", transactionResponse))
}

func (h *TransactionHandler) Update(c echo.Context) error {
	var midtrans MidtransCallbackRequest
	errBind := c.Bind(&midtrans)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}
	var midtransCore = MidtransCallbackReqestToCore(midtrans)
	err := h.transactionService.Update(midtransCore)
	if err != nil {
		if strings.Contains(err.Error(), "signature") {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500+" "+err.Error(), nil))
		}
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500+" "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", nil))
}

func (h *TransactionHandler) GetAllTicketDetail(c echo.Context) error {
	buyer_id, _ := middlewares.ExtractToken(c)
	result, err := h.transactionService.GetAllTicketDetail(buyer_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500+" "+err.Error(), nil))
	}
	var ticketDetailResponse = ListTicketDetailCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", ticketDetailResponse))
}
