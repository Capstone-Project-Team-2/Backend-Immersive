package buyers

import "github.com/labstack/echo/v4"

type BuyerCore struct {
	ID          uint
	Name        string
	PhoneNumber string
	Email       string
	Password    string
	Address     string
}

type BuyerDataInterface interface {
	Login(email, password string) (BuyerCore, error)
	SelectAll() ([]BuyerCore, error)
	Select(id string) (BuyerCore, error)
	Insert(input BuyerCore) error
	Update(input BuyerCore) error
	Delete(id string) error
}

type BuyerServiceInterface interface {
	Login(email, password string) (BuyerCore, string, error)
	Create(input BuyerCore) error
	GetAll() ([]BuyerCore, error)
	GetById(id string) (BuyerCore, error)
	UpdateById(input BuyerCore) error
	DeleteById(id string) error
}

type BuyerHandlerInterface interface {
	Login(c echo.Context) error
	Create(c echo.Context) error
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	UpdateById(c echo.Context) error
	DeleteById(c echo.Context) error
}
