package buyers

import (
	"mime/multipart"
)

type BuyerCore struct {
	ID          string
	Name        string `validate:"required"`
	PhoneNumber string
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	Address     string
	Avatar      string
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
	Register(input BuyerCore, file multipart.File) error
	ReadAll() ([]BuyerCore, error)
	Profile(id string) (BuyerCore, error)
	Edit(id string, input BuyerCore) error
	Deactive(id string) error
}
