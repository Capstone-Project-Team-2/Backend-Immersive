package handler

import (
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
	var transactionReq TransactionRequest

	errBind := c.Bind(&transactionReq)
	if errBind != nil {
		log.Error("handler - error on bind request")
		fmt.Println(errBind)
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400, nil))
	}

	newInput := TransactionRequestToCore(transactionReq)

	result, err := h.transactionService.Create(newInput)
	if err != nil {
		log.Error("handler-internal server error")
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	resultResp := TransactionCoreToResponse(result)
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "operation success", resultResp))

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
