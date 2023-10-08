package handler

import (
	"capstone-tickets/apps/middlewares"
	_eventHandler "capstone-tickets/features/events/handler"
	"capstone-tickets/features/transactions"
	"capstone-tickets/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

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
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	newInput := TransactionRequestToCore(transactionReq)

	err := h.transactionService.Create(newInput, buyer_id)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "operation success", nil))
}

func (h *TransactionHandler) GetById(c echo.Context) error {
	idParam := c.Param("transaction_id")
	buyer_id, role := middlewares.ExtractToken(c)
	if role != "Buyer" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	resultTrans, resultEvent, err := h.transactionService.Get(idParam, buyer_id)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var transactionResponse = TransactionCoreToResponse(resultTrans)
	var eventResponses = _eventHandler.EventCoreToResponse(resultEvent)
	data := map[string]any{
		"transaction": transactionResponse,
		"event":       eventResponses,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", data))
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
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500+" "+err.Error(), nil))
		}
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

func (h *TransactionHandler) GetAllPayment(c echo.Context) error {
	result, err := h.transactionService.GetAllPaymentMethod()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500+" "+err.Error(), nil))
	}
	var paymenResp = ListPaymentMethodCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", paymenResp))
}
