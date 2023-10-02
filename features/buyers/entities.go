package buyers

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type BuyerCore struct {
	ID          string
	Name        string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	Address     string
}

type BuyerDataInterface interface {
	Login(email, password string) (BuyerCore, string, error)
	Register(input BuyerCore, file multipart.File) error
	ReadAll() ([]BuyerCore, error)
	Profile(id string) (BuyerCore, error)
	Edit(input BuyerCore) error
	Deactive(id string) error
}

type BuyerServiceInterface interface {
	Login(email, password string) (BuyerCore, string, error)
	Register(input BuyerCore) error
	ReadAll() ([]BuyerCore, error)
	Profile(id string) (BuyerCore, error)
	Edit(id string, input BuyerCore) error
	Deactive(id string) error
}

type BuyerHandlerInterface interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	ReadAll(c echo.Context) error
	Profile(c echo.Context) error
	Edit(c echo.Context) error
	Deactive(c echo.Context) error
}
