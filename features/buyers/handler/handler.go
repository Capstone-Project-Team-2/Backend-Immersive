package handler

import (
	"capstone-tickets/features/buyers"

	"github.com/labstack/echo"
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

}
func (h *BuyerHandler) Register(c echo.Context) error {

}
